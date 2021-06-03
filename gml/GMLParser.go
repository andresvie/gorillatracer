package gml

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/light"
	"github.com/andresvie/gorillatracer/scene"
	"github.com/andresvie/gorillatracer/utils"
	"github.com/andresvie/gorillatracer/vector"
)

var sphereToken string = "sphere"
var pointOfLight string = "pointOfLight"
var tokens []string = []string{sphereToken, pointOfLight}

func ParseGML(reader io.Reader) (*scene.Scene, error) {
	fileScanner := bufio.NewScanner(reader)

	var lines []string
	var currentToken string = ""
	var tokenLine int = 0
	var line string
	var lineNumber int = 0
	var objects []geometry.Geometry = make([]geometry.Geometry, 0)
	var lights []light.Light = make([]light.Light, 0)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lineNumber++
		line = fileScanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		token, isToken := getToken(line, tokens)
		if isToken && !strings.Contains(line, "{") {
			return nil, errors.New(fmt.Sprintf("error in line %d expected open tag object %s", lineNumber, token))
		}
		if isToken {
			currentToken = token
			tokenLine = lineNumber
			continue
		}
		if strings.TrimSpace(line) == "}" && currentToken == "" {
			return nil, errors.New(fmt.Sprintf("error in line %d expected object declaration before close tag %s", lineNumber, token))
		}
		if strings.TrimSpace(line) == "}" && currentToken == sphereToken {
			sphere, err := parseSphere(lines, tokenLine)
			if err != nil {
				return nil, err
			}
			objects = append(objects, sphere)
			lines = nil
			continue
		}
		if strings.TrimSpace(line) == "}" && currentToken == pointOfLight {
			pointOfLight, err := parsePointOfLight(lines, tokenLine)
			if err != nil {
				return nil, err
			}
			lights = append(lights, pointOfLight)
			lines = nil
			continue
		}
		lines = append(lines, fileScanner.Text())

	}
	return &scene.Scene{Objects: objects, Lights: lights}, nil
}
func parsePointOfLight(lines []string, lineNumber int) (*light.PointOfLight, error) {
	var lightIntensity utils.REAL = 0
	var position *vector.Vector = nil
	var value interface{}
	var err error
	vectorTokens := []string{"point"}
	valueTokens := []string{"intensity"}
	for _, token := range lines {
		_, isToken := getToken(token, vectorTokens)
		_, isSingularToken := getToken(token, valueTokens)
		if isToken {
			_, value, err = parseVector(token, lineNumber)
			position = value.(*vector.Vector)
		}
		if isSingularToken && err == nil {
			_, value, err = parseReal(token, lineNumber)
			lightIntensity = value.(utils.REAL)
		}
		if err != nil {
			return nil, err
		}

	}
	if position == nil {
		return nil, errors.New(fmt.Sprintf("pointOfLight in the line %d point is required", lineNumber))
	}
	if lightIntensity == 0 {
		return nil, errors.New(fmt.Sprintf("pointOfLight in the line %d intensity is required", lineNumber))
	}
	return &light.PointOfLight{Point: position, Intensity: lightIntensity}, nil
}
func parseSphere(lines []string, lineNumber int) (*geometry.Sphere, error) {
	var color *vector.Vector = nil
	var position *vector.Vector = nil
	var radius utils.REAL = 0
	var specular utils.REAL = 0
	var reflection utils.REAL = 0
	var value interface{}
	var err error
	sphereVectorTokens := []string{"color", "center"}
	sphereValueTokens := []string{"radius", "specular", "reflection"}
	for _, token := range lines {
		sphereToken, isToken := getToken(token, sphereVectorTokens)
		sphereSingularToken, isSingularToken := getToken(token, sphereValueTokens)
		if isToken {
			_, value, err = parseVector(token, lineNumber)
		}
		if isSingularToken && err == nil {
			sphereToken = sphereSingularToken
			_, value, err = parseReal(token, lineNumber)
		}
		if err != nil {
			return nil, err
		}
		switch sphereToken {
		case "color":
			c, _ := value.(*vector.Vector)
			color = c
			break
		case "center":
			c, _ := value.(*vector.Vector)
			position = c
			break
		case "radius":
			c, _ := value.(utils.REAL)
			radius = c
			break
		case "reflection":
			c, _ := value.(utils.REAL)
			reflection = c
			break
		case "specular":
			c, _ := value.(utils.REAL)
			specular = c
			break
		}

	}

	if color == nil {
		return nil, errors.New(fmt.Sprintf("sphere in the line %d color is required", lineNumber))
	}
	if position == nil {
		return nil, errors.New(fmt.Sprintf("sphere in the line %d center is required", lineNumber))
	}
	if radius == 0 {
		return nil, errors.New(fmt.Sprintf("sphere in the line %d radius is required", lineNumber))
	}
	return &geometry.Sphere{Center: position, Color: color, Radius: radius, SpecularFactor: specular, ReflectionFactor: reflection}, nil
}

func parseReal(line string, lineNumber int) (string, utils.REAL, error) {
	pattern := regexp.MustCompile(`(?P<name>\w+) {0,}= {0,}(?P<value>\d+(\.\d+)?)`)
	matches := pattern.FindStringSubmatch(strings.TrimSpace(line))
	if !pattern.MatchString(line) {
		return "", 0, errors.New(fmt.Sprintf("a valid vector format is expected please review the object %d", lineNumber))
	}
	names := pattern.SubexpNames()
	positionMap := convertMatchToMap(matches, names)
	value := positionMap["value"]
	y, parseError := parseSingularValue(value, lineNumber)
	if parseError != nil {
		return "", 0, parseError
	}
	return positionMap["name"], y, parseError
}
func parseVector(line string, lineNumber int) (string, *vector.Vector, error) {
	pattern := regexp.MustCompile(`(?P<name>\w+) {0,}= {0,}\((?P<values>-?\d+(\.\d+)?,-?\d+(\.\d+)?,-?\d+(\.\d+)?)\)`)
	matches := pattern.FindStringSubmatch(strings.TrimSpace(line))
	if !pattern.MatchString(line) {
		return "", nil, errors.New(fmt.Sprintf("a valid vector format is expected please review the object %d", lineNumber))
	}
	names := pattern.SubexpNames()
	positionMap := convertMatchToMap(matches, names)
	values := strings.Split(positionMap["values"], ",")
	x, parseError := parseSingularValue(values[0], lineNumber)
	if parseError != nil {
		return "", nil, parseError
	}
	y, parseError := parseSingularValue(values[1], lineNumber)
	if parseError != nil {
		return "", nil, parseError
	}
	z, parseError := parseSingularValue(values[2], lineNumber)
	if parseError != nil {
		return "", nil, parseError
	}
	return positionMap["name"], &vector.Vector{X: x, Y: y, Z: z, W: 0}, nil
}

func parseSingularValue(textValue string, lineNumber int) (utils.REAL, error) {
	x, parseError := strconv.ParseFloat(textValue, 64)
	if parseError != nil {
		return 0, errors.New(fmt.Sprintf("a valid position values are expected please review the object %d", lineNumber))
	}
	return utils.REAL(x), nil
}

func convertMatchToMap(matches, names []string) map[string]string {
	var m map[string]string = make(map[string]string)
	for i, match := range matches {
		if i != 0 {
			m[names[i]] = match
		}
	}
	return m
}

func getToken(line string, tokens []string) (string, bool) {
	for _, token := range tokens {
		if strings.Contains(line, token) {
			return token, true
		}
	}

	return "", false
}
