package options

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

func TestNewSchool(t *testing.T) {
	s := NewSchool()
	assert.NotNil(t, s)
	assert.Equal(t, defaultStudents, s.students)
	assert.Equal(t, defaultTeachers, s.teachers)
}

func TestNewSchool_WithOptions(t *testing.T) {
	num := 1
	s1 := NewSchool(SetStudents(num))
	assert.NotNil(t, s1)
	s2 := NewSchool(SetTeachers(num))
	assert.NotNil(t, s2)
	s3 := NewSchool(SetStudents(num), SetTeachers(num))
	assert.NotNil(t, s3)
}

func TestSetStudents(t *testing.T) {
	studentNum := gofakeit.Number(1, 100)
	teacherNum := gofakeit.Number(1, 100)
	s := NewSchool(SetStudents(studentNum), SetTeachers(teacherNum))
	assert.Equal(t, studentNum, s.students)
	assert.Equal(t, teacherNum, s.teachers)
}
