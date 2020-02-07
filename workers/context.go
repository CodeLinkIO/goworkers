package workers

type WorkerContext struct {
	ID string
}

type ContextKey string

const WorkerContextKey = ContextKey("WorkerContext")
