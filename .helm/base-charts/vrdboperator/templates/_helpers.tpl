
{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "vrdboperator.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}


{{/*
Default annotations for jnnkrdb github projects.
*/}}
{{- define "vrdboperator.defaultAnnotations" -}}
jnnkrdb.de/source: "github.com/jnnkrdb/vaultrdb"
{{- end }}

{{/*
Common labels 
*/}}
{{- define "vrdboperator.labels" -}}
jnnkrdb.de/chart: {{ include "vrdboperator.chart" . }}
helm.sh/chart: {{ include "vrdboperator.chart" . }}
{{ include "vrdboperator.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels 
*/}}
{{- define "vrdboperator.selectorLabels" -}}
jnnkrdb.de/service: vrdboperator
jnnkrdb.de/instance: {{ .Release.Name }}
{{- end }}


