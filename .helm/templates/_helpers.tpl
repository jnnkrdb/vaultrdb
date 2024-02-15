{{/*
Expand the name of the chart.
*/}}
{{- define "helmbase.name" -}}
{{- default .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "helmbase.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "helmbase.labels" -}}
helm.sh/chart: {{ include "helmbase.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service | quote }}
meta.helm.sh/release-name: {{ .Release.Name }}
meta.helm.sh/release-namespace: {{ .Values.namespace }}
{{ include "helmbase.selectorLabels" . }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "helmbase.selectorLabels" -}}
jnnkrdb.de/service: "vaultrdb"
app.kubernetes.io/name: {{ include "helmbase.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
