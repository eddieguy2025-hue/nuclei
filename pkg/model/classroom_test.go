package model

import (
	"testing"

	"github.com/projectdiscovery/nuclei/v3/pkg/utils/json"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestClassroomJsonMarshal(t *testing.T) {
	classroom := Classroom{
		Name:     "Computer Lab 101",
		Capacity: 30,
		Location: "Building A, Floor 2",
	}

	result, err := json.Marshal(&classroom)
	require.Nil(t, err)

	expected := `{"name":"Computer Lab 101","capacity":30,"location":"Building A, Floor 2"}`
	require.Equal(t, expected, string(result))
}

func TestClassroomYamlMarshal(t *testing.T) {
	classroom := Classroom{
		Name:     "Computer Lab 101",
		Capacity: 30,
		Location: "Building A, Floor 2",
	}

	result, err := yaml.Marshal(&classroom)
	require.Nil(t, err)

	expected := `name: Computer Lab 101
capacity: 30
location: Building A, Floor 2
`
	require.Equal(t, expected, string(result))
}

func TestClassroomYamlUnmarshal(t *testing.T) {
	yamlPayload := `
name: Science Lab 202
capacity: 25
location: Building B, Floor 1
`
	classroom := Classroom{}
	err := yaml.Unmarshal([]byte(yamlPayload), &classroom)
	require.Nil(t, err)
	require.Equal(t, "Science Lab 202", classroom.Name)
	require.Equal(t, 25, classroom.Capacity)
	require.Equal(t, "Building B, Floor 1", classroom.Location)
}

func TestNewClassroom(t *testing.T) {
	classroom := NewClassroom("Math Room 303", 40, "Building C, Floor 3")
	require.NotNil(t, classroom)
	require.Equal(t, "Math Room 303", classroom.Name)
	require.Equal(t, 40, classroom.Capacity)
	require.Equal(t, "Building C, Floor 3", classroom.Location)
}
