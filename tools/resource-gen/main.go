package main

import (
	"bytes"
	"flag"
	"go/format"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
)

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

import (
	_ "github.com/apache/dubbo-kubernetes/api/mesh/v1alpha1"
	_ "github.com/apache/dubbo-kubernetes/api/system/v1alpha1"
	. "github.com/apache/dubbo-kubernetes/tools/resource-gen/genutils"
)

// CustomResourceTemplate for creating a Kubernetes CRD to wrap a Dubbo resource.
var CustomResourceTemplate = template.Must(template.New("custom-resource").Parse(`
// Generated by tools/resource-gen
// Run "make generate" to update this file.

{{ $pkg := printf "%s_proto" .Package }}
{{ $tk := "` + "`" + `" }}

// nolint:whitespace
package v1alpha1

import (
	"fmt"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	{{ $pkg }} "github.com/apache/dubbo-kubernetes/api/{{ .Package }}/v1alpha1"
	core_model "github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/model"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/registry"
	util_proto "github.com/apache/dubbo-kubernetes/pkg/util/proto"
)

{{range .Resources}}
{{- if not .SkipKubernetesWrappers }}

// +kubebuilder:object:root=true
{{- if .ScopeNamespace }}
// +kubebuilder:resource:categories=dubbo,scope=Namespaced
{{- else }}
// +kubebuilder:resource:categories=dubbo,scope=Cluster
{{- end}}
{{- range .AdditionalPrinterColumns }}
// +kubebuilder:printcolumn:{{ . }}
{{- end}}
type {{.ResourceType}} struct {
	metav1.TypeMeta   {{ $tk }}json:",inline"{{ $tk }}
	metav1.ObjectMeta {{ $tk }}json:"metadata,omitempty"{{ $tk }}

    // Mesh is the name of the dubbo mesh this resource belongs to.
	// It may be omitted for cluster-scoped resources.
	//
    // +kubebuilder:validation:Optional
	Mesh string {{ $tk }}json:"mesh,omitempty"{{ $tk }}

{{- if eq .ResourceType "DataplaneInsight" }}
	// Status is the status the dubbo resource.
    // +kubebuilder:validation:Optional
	Status   *apiextensionsv1.JSON {{ $tk }}json:"status,omitempty"{{ $tk }}
{{- else}}
	// Spec is the specification of the Dubbo {{ .ProtoType }} resource.
    // +kubebuilder:validation:Optional
	Spec   *apiextensionsv1.JSON {{ $tk }}json:"spec,omitempty"{{ $tk }}
{{- end}}
}

// +kubebuilder:object:root=true
{{- if .ScopeNamespace }}
// +kubebuilder:resource:scope=Cluster
{{- else }}
// +kubebuilder:resource:scope=Namespaced
{{- end}}
type {{.ResourceType}}List struct {
	metav1.TypeMeta {{ $tk }}json:",inline"{{ $tk }}
	metav1.ListMeta {{ $tk }}json:"metadata,omitempty"{{ $tk }}
	Items           []{{.ResourceType}} {{ $tk }}json:"items"{{ $tk }}
}

{{- if not .SkipRegistration}}
func init() {
	SchemeBuilder.Register(&{{.ResourceType}}{}, &{{.ResourceType}}List{})
}
{{- end}}

func (cb *{{.ResourceType}}) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *{{.ResourceType}}) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *{{.ResourceType}}) GetMesh() string {
	return cb.Mesh
}

func (cb *{{.ResourceType}}) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *{{.ResourceType}}) GetSpec() (core_model.ResourceSpec, error) {
{{- if eq .ResourceType "DataplaneInsight" }}
	spec := cb.Status
{{- else}}
	spec := cb.Spec
{{- end}}
	m := {{$pkg}}.{{.ProtoType}}{}

    if spec == nil || len(spec.Raw) == 0 {
		return &m, nil
	}

	err := util_proto.FromJSON(spec.Raw, &m)
	return &m, err
}

func (cb *{{.ResourceType}}) SetSpec(spec core_model.ResourceSpec) {
	if spec == nil {
{{- if eq .ResourceType "DataplaneInsight" }}
		cb.Status = nil
{{- else }}
		cb.Spec = nil
{{- end }}
		return
	}

	s, ok := spec.(*{{$pkg}}.{{.ProtoType}}); 
	if !ok {
		panic(fmt.Sprintf("unexpected protobuf message type %T", spec))
	}

{{ if eq .ResourceType "DataplaneInsight" }}
	cb.Status = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
{{- else}}
	cb.Spec = &apiextensionsv1.JSON{Raw: util_proto.MustMarshalJSON(s)}
{{- end}}
}

func (cb *{{.ResourceType}}) Scope() model.Scope {
{{- if .ScopeNamespace }}
	return model.ScopeNamespace
{{- else }}
	return model.ScopeCluster
{{- end }}
}

func (l *{{.ResourceType}}List) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

{{if not .SkipRegistration}}
func init() {
	registry.RegisterObjectType(&{{ $pkg }}.{{.ProtoType}}{}, &{{.ResourceType}}{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "{{.ResourceType}}",
		},
	})
	registry.RegisterListType(&{{ $pkg }}.{{.ProtoType}}{}, &{{.ResourceType}}List{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "{{.ResourceType}}List",
		},
	})
}
{{- end }} {{/* .SkipRegistration */}}
{{- end }} {{/* .SkipKubernetesWrappers */}}
{{- end }} {{/* Resources */}}
`))

// ResourceTemplate for creating a Dubbo resource.
var ResourceTemplate = template.Must(template.New("resource").Funcs(map[string]any{"hasSuffix": strings.HasSuffix, "trimSuffix": strings.TrimSuffix}).Parse(`
// Generated by tools/resource-gen.
// Run "make generate" to update this file.

{{ $pkg := printf "%s_proto" .Package }}

// nolint:whitespace
package {{.Package}}

import (
	"fmt"

	{{$pkg}} "github.com/apache/dubbo-kubernetes/api/{{.Package}}/v1alpha1"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/registry"
)

{{range .Resources}}
{{ $baseType := trimSuffix (trimSuffix .ResourceType "Overview") "Insight" }}
const (
	{{.ResourceType}}Type model.ResourceType = "{{.ResourceType}}"
)

var _ model.Resource = &{{.ResourceName}}{}

type {{.ResourceName}} struct {
	Meta model.ResourceMeta
	Spec *{{$pkg}}.{{.ProtoType}}
}

func New{{.ResourceName}}() *{{.ResourceName}} {
	return &{{.ResourceName}}{
		Spec: &{{$pkg}}.{{.ProtoType}}{},
	}
}

func (t *{{.ResourceName}}) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *{{.ResourceName}}) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *{{.ResourceName}}) GetSpec() model.ResourceSpec {
	return t.Spec
}

{{with $in := .}}
{{range .Selectors}}
func (t *{{$in.ResourceName}}) {{.}}() []*{{$pkg}}.Selector {
	return t.Spec.Get{{.}}()
}
{{end}}
{{end}}

func (t *{{.ResourceName}}) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*{{$pkg}}.{{.ProtoType}})
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &{{$pkg}}.{{.ProtoType}}{}
		} else  {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *{{.ResourceName}}) Descriptor() model.ResourceTypeDescriptor {
	return {{.ResourceName}}TypeDescriptor 
}
{{- if and (hasSuffix .ResourceType "Overview") (ne $baseType "Service") }}

func (t *{{.ResourceName}}) SetOverviewSpec(resource model.Resource, insight model.Resource) error {
	t.SetMeta(resource.GetMeta())
	overview := &{{$pkg}}.{{.ProtoType}}{
		{{$baseType}}: resource.GetSpec().(*{{$pkg}}.{{$baseType}}),
	}
	if insight != nil {
		ins, ok := insight.GetSpec().(*{{$pkg}}.{{$baseType}}Insight)
		if !ok {
			return errors.New("failed to convert to insight type '{{$baseType}}Insight'")
		}
		overview.{{$baseType}}Insight = ins
	}
	return t.SetSpec(overview)
}
{{- end }}

var _ model.ResourceList = &{{.ResourceName}}List{}

type {{.ResourceName}}List struct {
	Items      []*{{.ResourceName}}
	Pagination model.Pagination
}

func (l *{{.ResourceName}}List) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *{{.ResourceName}}List) GetItemType() model.ResourceType {
	return {{.ResourceType}}Type
}

func (l *{{.ResourceName}}List) NewItem() model.Resource {
	return New{{.ResourceName}}()
}

func (l *{{.ResourceName}}List) AddItem(r model.Resource) error {
	if trr, ok := r.(*{{.ResourceName}}); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*{{.ResourceName}})(nil), r)
	}
}

func (l *{{.ResourceName}}List) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *{{.ResourceName}}List) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var {{.ResourceName}}TypeDescriptor = model.ResourceTypeDescriptor{
		Name: {{.ResourceType}}Type,
		Resource: New{{.ResourceName}}(),
		ResourceList: &{{.ResourceName}}List{},
		ReadOnly: {{.WsReadOnly}},
		AdminOnly: {{.WsAdminOnly}},
		Scope: {{if .Global}}model.ScopeGlobal{{else}}model.ScopeMesh{{end}},
		{{- if ne .DdsDirection ""}}
		DDSFlags: {{.DdsDirection}},
		{{- end}}
		WsPath: "{{.WsPath}}",
		DubboctlArg: "{{.DubboctlSingular}}",
		DubboctlListArg: "{{.DubboctlPlural}}",
		AllowToInspect: {{.AllowToInspect}},
		IsPolicy: {{.IsPolicy}},
		SingularDisplayName: "{{.SingularDisplayName}}",
		PluralDisplayName: "{{.PluralDisplayName}}",
		IsExperimental: {{.IsExperimental}},
	}

{{- if not .SkipRegistration}}
func init() {
	registry.RegisterType({{.ResourceName}}TypeDescriptor)
}
{{- end}}
{{end}}
`))

// ProtoMessageFunc ...
type ProtoMessageFunc func(protoreflect.MessageType) bool

// OnDubboResourceMessage ...
func OnDubboResourceMessage(pkg string, f ProtoMessageFunc) ProtoMessageFunc {
	return func(m protoreflect.MessageType) bool {
		r := DubboResourceForMessage(m.Descriptor())
		if r == nil {
			return true
		}

		if r.Package == pkg {
			return f(m)
		}

		return true
	}
}

func main() {
	var gen string
	var pkg string

	flag.StringVar(&gen, "generator", "", "the type of generator to run options: (type,crd)")
	flag.StringVar(&pkg, "package", "", "the name of the package to generate: (mesh, system)")

	flag.Parse()

	switch pkg {
	case "mesh", "system":
	default:
		log.Fatalf("package %s is not supported", pkg)
	}

	var types []protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(
		OnDubboResourceMessage(pkg, func(m protoreflect.MessageType) bool {
			types = append(types, m)
			return true
		}))

	// Sort by name so the output is deterministic.
	sort.Slice(types, func(i, j int) bool {
		return types[i].Descriptor().FullName() < types[j].Descriptor().FullName()
	})

	var resources []ResourceInfo
	for _, t := range types {
		resourceInfo := ToResourceInfo(t.Descriptor())
		resources = append(resources, resourceInfo)
	}

	var generatorTemplate *template.Template

	switch gen {
	case "type":
		generatorTemplate = ResourceTemplate
	case "crd":
		generatorTemplate = CustomResourceTemplate
	default:
		log.Fatalf("%s is not a valid generator option\n", gen)
	}

	outBuf := bytes.Buffer{}
	if err := generatorTemplate.Execute(&outBuf, struct {
		Package   string
		Resources []ResourceInfo
	}{
		Package:   pkg,
		Resources: resources,
	}); err != nil {
		log.Fatalf("template error: %s", err)
	}

	out, err := format.Source(outBuf.Bytes())
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	if _, err := os.Stdout.Write(out); err != nil {
		log.Fatalf("%s\n", err)
	}
}