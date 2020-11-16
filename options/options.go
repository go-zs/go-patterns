package options

type (
	school struct {
		students int
		teachers int
	}

	SchoolOption interface {
		apply(*school)
	}
)

const (
	defaultStudents = 100
	defaultTeachers = 500
)

func NewSchool(options ...SchoolOption) *school {
	s := &school{
		students: defaultStudents,
		teachers: defaultTeachers,
	}
	for _, o := range options {
		o.apply(s)
	}

	return s
}

func SetStudents(num int) SchoolOption {
	return studentOption(num)
}

func SetTeachers(num int) SchoolOption {
	return teacherOption(num)
}

type studentOption int

func (o studentOption) apply(s *school) {
	s.students = int(o)
}

type teacherOption int

func (o teacherOption) apply(s *school) {
	s.teachers = int(o)
}
