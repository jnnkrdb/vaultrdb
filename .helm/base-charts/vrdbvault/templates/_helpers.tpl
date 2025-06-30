
{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "vrdbvault.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}


{{/*
Default annotations for jnnkrdb github projects.
*/}}
{{- define "vrdbvault.defaultAnnotations" -}}
jnnkrdb.de/source: "github.com/jnnkrdb/vaultrdb"
{{- end }}

{{/*
Common labels 
*/}}
{{- define "vrdbvault.labels" -}}
jnnkrdb.de/chart: {{ include "vrdbvault.chart" . }}
helm.sh/chart: {{ include "vrdbvault.chart" . }}
{{ include "vrdbvault.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels 
*/}}
{{- define "vrdbvault.selectorLabels" -}}
jnnkrdb.de/service: vrdbvault
jnnkrdb.de/instance: {{ .Release.Name }}
{{- end }}


