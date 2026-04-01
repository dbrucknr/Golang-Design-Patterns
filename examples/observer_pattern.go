package examples

func observerSample() {
	s := &SimpleSubject{}

	o := NewSimpleObserver(s)
	s.SetValue(1)
	o.DisplayValue()

	s.SetValue(2)
	o.DisplayValue()
}

// The 'One' in the One to Many Relationship
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type SimpleSubject struct {
	value     int
	observers []Observer
}

// / Add an observer to the subject's list of observers.
func (s *SimpleSubject) RegisterObserver(o Observer) {
	s.observers = append(s.observers, o)
}

// / Remove an observer from the subject's list of observers.
func (s *SimpleSubject) RemoveObserver(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *SimpleSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.PullUpdate()
		observer.PushUpdate(s.value)
	}
}

func (s *SimpleSubject) SetValue(value int) {
	s.value = value
	s.NotifyObservers()
}

// The 'Many' in the One to Many Relationship
type Observer interface {
	PullUpdate()
	PushUpdate(value int)
}

type SimpleObserver struct {
	value   int
	subject *SimpleSubject
}

// The 'Pull' model: the subject says 'data changed', and the observer fetches the value from the subject.
func (o *SimpleObserver) PullUpdate() {
	println("Observer pulled an update")
	o.value = o.subject.value
}

// The 'Push' model: the subject sends the updated value directly to the observer.
func (o *SimpleObserver) PushUpdate(value int) {
	println("Observer pushed an update")
	o.value = value
}

func (o *SimpleObserver) DisplayValue() int {
	println(o.value)
	return o.value
}

func NewSimpleObserver(subject *SimpleSubject) *SimpleObserver {
	o := &SimpleObserver{subject: subject}
	subject.RegisterObserver(o)
	return o
}
