{{- define "revision" -}}
{{- .Values.global.revision | replace "." "-" -}}
{{- end -}}

{{- define "namespaced-revision" -}}
{{ $revision := (include "revision" .) }}
{{ if $revision eq "default" -}}
{{- $revision -}}
{{- else -}}
{{- printf "%s.%s" $revision .Release.Namespace -}}
{{- end -}}
{{- end -}}

{{- define "name-with-revision" -}}
{{- if .context.Values.global.revision -}}
{{- printf "%s-%s" .name (include "revision" .context) -}}
{{- else -}}
{{- .name -}}
{{- end -}}
{{- end -}}

{{- define "name-with-namespaced-revision" -}}
{{- if .context.Values.global.revision -}}
{{- printf "%s-%s" (include "name-with-revision" .) .context.Release.Namespace -}}
{{- else -}}
{{- printf "%s-%s" .name .context.Release.Namespace -}}
{{- end -}}
{{- end -}}
