package view

import (
	"github.com/derailed/k9s/internal/client"
	"github.com/derailed/k9s/internal/render"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

// Job represents a job viewer.
type Job struct {
	ResourceViewer
}

// NewJob returns a new viewer.
func NewJob(gvr client.GVR) ResourceViewer {
	j := Job{ResourceViewer: NewLogsExtender(NewBrowser(gvr), nil)}
	j.GetTable().SetEnterFn(j.showPods)
	j.GetTable().SetColorerFn(render.Job{}.ColorerFunc())

	return &j
}

func (*Job) showPods(app *App, _, gvr, path string) {
	o, err := app.factory.Get(gvr, path, labels.Everything())
	if err != nil {
		app.Flash().Err(err)
		return
	}

	var job batchv1.Job
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.(*unstructured.Unstructured).Object, &job)
	if err != nil {
		app.Flash().Err(err)
		return
	}

	showPodsFromSelector(app, path, job.Spec.Selector)
}
