package go_tests

import (
	"testing"

	"github.com/keptn/keptn/configuration-service/models"
	"github.com/stretchr/testify/require"
)

func Test_RestoreServiceBasic(t *testing.T) {
	projectName := "resource-service-test-project"
	nonExistingProjectName := "non_existing_project"
	nonExistingStageName := "non_existing_stage"
	nonExistingServiceName := "non_existing_service"

	//gitUser := getGiteaUser()
	//gitToken, err := getGiteaToken()
	//require.Nil(t, err)
	//gitRemoteURI := fmt.Sprintf("http://gitea-http:3000/%s/%s", gitUser, projectName)

	resourceContent := "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPGptZXRlclRlc3RQbGFuIHZlcnNpb249IjEuMiIgcHJvcGVydGllcz0iNS4wIiBqbWV0ZXI9IjUuNCI+CiAgPGhhc2hUcmVlPgogICAgPFRlc3RQbGFuIGd1aWNsYXNzPSJUZXN0UGxhbkd1aSIgdGVzdGNsYXNzPSJUZXN0UGxhbiIgdGVzdG5hbWU9IlRlc3QgUGxhbiIgZW5hYmxlZD0idHJ1ZSI+CiAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IlRlc3RQbGFuLmNvbW1lbnRzIj48L3N0cmluZ1Byb3A+CiAgICAgIDxib29sUHJvcCBuYW1lPSJUZXN0UGxhbi5mdW5jdGlvbmFsX21vZGUiPmZhbHNlPC9ib29sUHJvcD4KICAgICAgPGJvb2xQcm9wIG5hbWU9IlRlc3RQbGFuLnNlcmlhbGl6ZV90aHJlYWRncm91cHMiPmZhbHNlPC9ib29sUHJvcD4KICAgICAgPGVsZW1lbnRQcm9wIG5hbWU9IlRlc3RQbGFuLnVzZXJfZGVmaW5lZF92YXJpYWJsZXMiIGVsZW1lbnRUeXBlPSJBcmd1bWVudHMiIGd1aWNsYXNzPSJBcmd1bWVudHNQYW5lbCIgdGVzdGNsYXNzPSJBcmd1bWVudHMiIHRlc3RuYW1lPSJVc2VyIERlZmluZWQgVmFyaWFibGVzIiBlbmFibGVkPSJ0cnVlIj4KICAgICAgICA8Y29sbGVjdGlvblByb3AgbmFtZT0iQXJndW1lbnRzLmFyZ3VtZW50cyI+CiAgICAgICAgICA8ZWxlbWVudFByb3AgbmFtZT0iU0VSVkVSX1VSTCIgZWxlbWVudFR5cGU9IkFyZ3VtZW50Ij4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQubmFtZSI+U0VSVkVSX1VSTDwvc3RyaW5nUHJvcD4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQudmFsdWUiPmVjMi01NC0xNjQtMTY0LTEyMS5jb21wdXRlLTEuYW1hem9uYXdzLmNvbTwvc3RyaW5nUHJvcD4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQubWV0YWRhdGEiPj08L3N0cmluZ1Byb3A+CiAgICAgICAgICA8L2VsZW1lbnRQcm9wPgogICAgICAgICAgPGVsZW1lbnRQcm9wIG5hbWU9IkRlZmF1bHRUaGlua1RpbWUiIGVsZW1lbnRUeXBlPSJBcmd1bWVudCI+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50Lm5hbWUiPkRlZmF1bHRUaGlua1RpbWU8L3N0cmluZ1Byb3A+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50LnZhbHVlIj4yNTA8L3N0cmluZ1Byb3A+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50Lm1ldGFkYXRhIj49PC9zdHJpbmdQcm9wPgogICAgICAgICAgPC9lbGVtZW50UHJvcD4KICAgICAgICAgIDxlbGVtZW50UHJvcCBuYW1lPSJEVF9MVE4iIGVsZW1lbnRUeXBlPSJBcmd1bWVudCI+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50Lm5hbWUiPkRUX0xUTjwvc3RyaW5nUHJvcD4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQudmFsdWUiPlRlc3RKdW5lMDM8L3N0cmluZ1Byb3A+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50Lm1ldGFkYXRhIj49PC9zdHJpbmdQcm9wPgogICAgICAgICAgPC9lbGVtZW50UHJvcD4KICAgICAgICAgIDxlbGVtZW50UHJvcCBuYW1lPSJTRVJWRVJfUE9SVCIgZWxlbWVudFR5cGU9IkFyZ3VtZW50Ij4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQubmFtZSI+U0VSVkVSX1BPUlQ8L3N0cmluZ1Byb3A+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50LnZhbHVlIj44MDwvc3RyaW5nUHJvcD4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQubWV0YWRhdGEiPj08L3N0cmluZ1Byb3A+CiAgICAgICAgICA8L2VsZW1lbnRQcm9wPgogICAgICAgICAgPGVsZW1lbnRQcm9wIG5hbWU9Ikxvb3BDb3VudCIgZWxlbWVudFR5cGU9IkFyZ3VtZW50Ij4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQubmFtZSI+TG9vcENvdW50PC9zdHJpbmdQcm9wPgogICAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJBcmd1bWVudC52YWx1ZSI+MTAwMDwvc3RyaW5nUHJvcD4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQubWV0YWRhdGEiPj08L3N0cmluZ1Byb3A+CiAgICAgICAgICA8L2VsZW1lbnRQcm9wPgogICAgICAgICAgPGVsZW1lbnRQcm9wIG5hbWU9IkNIRUNLX1BBVEgiIGVsZW1lbnRUeXBlPSJBcmd1bWVudCI+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50Lm5hbWUiPkNIRUNLX1BBVEg8L3N0cmluZ1Byb3A+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50LnZhbHVlIj4vPC9zdHJpbmdQcm9wPgogICAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJBcmd1bWVudC5tZXRhZGF0YSI+PTwvc3RyaW5nUHJvcD4KICAgICAgICAgIDwvZWxlbWVudFByb3A+CiAgICAgICAgICA8ZWxlbWVudFByb3AgbmFtZT0iUFJPVE9DT0wiIGVsZW1lbnRUeXBlPSJBcmd1bWVudCI+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50Lm5hbWUiPlBST1RPQ09MPC9zdHJpbmdQcm9wPgogICAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJBcmd1bWVudC52YWx1ZSI+aHR0cDwvc3RyaW5nUHJvcD4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQubWV0YWRhdGEiPj08L3N0cmluZ1Byb3A+CiAgICAgICAgICA8L2VsZW1lbnRQcm9wPgogICAgICAgICAgPGVsZW1lbnRQcm9wIG5hbWU9IlZVQ291bnQiIGVsZW1lbnRUeXBlPSJBcmd1bWVudCI+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50Lm5hbWUiPlZVQ291bnQ8L3N0cmluZ1Byb3A+CiAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkFyZ3VtZW50LnZhbHVlIj4xMDwvc3RyaW5nUHJvcD4KICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQXJndW1lbnQubWV0YWRhdGEiPj08L3N0cmluZ1Byb3A+CiAgICAgICAgICA8L2VsZW1lbnRQcm9wPgogICAgICAgIDwvY29sbGVjdGlvblByb3A+CiAgICAgIDwvZWxlbWVudFByb3A+CiAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IlRlc3RQbGFuLnVzZXJfZGVmaW5lX2NsYXNzcGF0aCI+PC9zdHJpbmdQcm9wPgogICAgPC9UZXN0UGxhbj4KICAgIDxoYXNoVHJlZT4KICAgICAgPFRocmVhZEdyb3VwIGd1aWNsYXNzPSJUaHJlYWRHcm91cEd1aSIgdGVzdGNsYXNzPSJUaHJlYWRHcm91cCIgdGVzdG5hbWU9IlRocmVhZCBHcm91cCIgZW5hYmxlZD0idHJ1ZSI+CiAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iVGhyZWFkR3JvdXAub25fc2FtcGxlX2Vycm9yIj5jb250aW51ZTwvc3RyaW5nUHJvcD4KICAgICAgICA8ZWxlbWVudFByb3AgbmFtZT0iVGhyZWFkR3JvdXAubWFpbl9jb250cm9sbGVyIiBlbGVtZW50VHlwZT0iTG9vcENvbnRyb2xsZXIiIGd1aWNsYXNzPSJMb29wQ29udHJvbFBhbmVsIiB0ZXN0Y2xhc3M9Ikxvb3BDb250cm9sbGVyIiB0ZXN0bmFtZT0iTG9vcCBDb250cm9sbGVyIiBlbmFibGVkPSJ0cnVlIj4KICAgICAgICAgIDxib29sUHJvcCBuYW1lPSJMb29wQ29udHJvbGxlci5jb250aW51ZV9mb3JldmVyIj5mYWxzZTwvYm9vbFByb3A+CiAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJMb29wQ29udHJvbGxlci5sb29wcyI+JHtfX1AoTG9vcENvdW50LCR7TG9vcENvdW50fSl9PC9zdHJpbmdQcm9wPgogICAgICAgIDwvZWxlbWVudFByb3A+CiAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iVGhyZWFkR3JvdXAubnVtX3RocmVhZHMiPiR7X19QKFZVQ291bnQsJHtWVUNvdW50fSl9PC9zdHJpbmdQcm9wPgogICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IlRocmVhZEdyb3VwLnJhbXBfdGltZSI+MTwvc3RyaW5nUHJvcD4KICAgICAgICA8bG9uZ1Byb3AgbmFtZT0iVGhyZWFkR3JvdXAuc3RhcnRfdGltZSI+MTUzNjA2NDUxNzAwMDwvbG9uZ1Byb3A+CiAgICAgICAgPGxvbmdQcm9wIG5hbWU9IlRocmVhZEdyb3VwLmVuZF90aW1lIj4xNTM2MDY0NTE3MDAwPC9sb25nUHJvcD4KICAgICAgICA8Ym9vbFByb3AgbmFtZT0iVGhyZWFkR3JvdXAuc2NoZWR1bGVyIj5mYWxzZTwvYm9vbFByb3A+CiAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iVGhyZWFkR3JvdXAuZHVyYXRpb24iPjwvc3RyaW5nUHJvcD4KICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJUaHJlYWRHcm91cC5kZWxheSI+PC9zdHJpbmdQcm9wPgogICAgICAgIDxib29sUHJvcCBuYW1lPSJUaHJlYWRHcm91cC5zYW1lX3VzZXJfb25fbmV4dF9pdGVyYXRpb24iPnRydWU8L2Jvb2xQcm9wPgogICAgICA8L1RocmVhZEdyb3VwPgogICAgICA8aGFzaFRyZWU+CiAgICAgICAgPENvb2tpZU1hbmFnZXIgZ3VpY2xhc3M9IkNvb2tpZVBhbmVsIiB0ZXN0Y2xhc3M9IkNvb2tpZU1hbmFnZXIiIHRlc3RuYW1lPSJIVFRQIENvb2tpZSBNYW5hZ2VyIiBlbmFibGVkPSJ0cnVlIj4KICAgICAgICAgIDxjb2xsZWN0aW9uUHJvcCBuYW1lPSJDb29raWVNYW5hZ2VyLmNvb2tpZXMiLz4KICAgICAgICAgIDxib29sUHJvcCBuYW1lPSJDb29raWVNYW5hZ2VyLmNsZWFyRWFjaEl0ZXJhdGlvbiI+ZmFsc2U8L2Jvb2xQcm9wPgogICAgICAgICAgPGJvb2xQcm9wIG5hbWU9IkNvb2tpZU1hbmFnZXIuY29udHJvbGxlZEJ5VGhyZWFkR3JvdXAiPmZhbHNlPC9ib29sUHJvcD4KICAgICAgICA8L0Nvb2tpZU1hbmFnZXI+CiAgICAgICAgPGhhc2hUcmVlLz4KICAgICAgICA8SGVhZGVyTWFuYWdlciBndWljbGFzcz0iSGVhZGVyUGFuZWwiIHRlc3RjbGFzcz0iSGVhZGVyTWFuYWdlciIgdGVzdG5hbWU9IkhUVFAgSGVhZGVyIE1hbmFnZXIiIGVuYWJsZWQ9InRydWUiPgogICAgICAgICAgPGNvbGxlY3Rpb25Qcm9wIG5hbWU9IkhlYWRlck1hbmFnZXIuaGVhZGVycyI+CiAgICAgICAgICAgIDxlbGVtZW50UHJvcCBuYW1lPSIiIGVsZW1lbnRUeXBlPSJIZWFkZXIiPgogICAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkhlYWRlci5uYW1lIj5DYWNoZS1Db250cm9sPC9zdHJpbmdQcm9wPgogICAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkhlYWRlci52YWx1ZSI+bm8tY2FjaGU8L3N0cmluZ1Byb3A+CiAgICAgICAgICAgIDwvZWxlbWVudFByb3A+CiAgICAgICAgICAgIDxlbGVtZW50UHJvcCBuYW1lPSIiIGVsZW1lbnRUeXBlPSJIZWFkZXIiPgogICAgICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkhlYWRlci5uYW1lIj5Db250ZW50LVR5cGU8L3N0cmluZ1Byb3A+CiAgICAgICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iSGVhZGVyLnZhbHVlIj5hcHBsaWNhdGlvbi9qc29uPC9zdHJpbmdQcm9wPgogICAgICAgICAgICA8L2VsZW1lbnRQcm9wPgogICAgICAgICAgICA8ZWxlbWVudFByb3AgbmFtZT0iIiBlbGVtZW50VHlwZT0iSGVhZGVyIj4KICAgICAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJIZWFkZXIubmFtZSI+anNvbjwvc3RyaW5nUHJvcD4KICAgICAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJIZWFkZXIudmFsdWUiPnRydWU8L3N0cmluZ1Byb3A+CiAgICAgICAgICAgIDwvZWxlbWVudFByb3A+CiAgICAgICAgICA8L2NvbGxlY3Rpb25Qcm9wPgogICAgICAgIDwvSGVhZGVyTWFuYWdlcj4KICAgICAgICA8aGFzaFRyZWUvPgogICAgICAgIDxCZWFuU2hlbGxQcmVQcm9jZXNzb3IgZ3VpY2xhc3M9IlRlc3RCZWFuR1VJIiB0ZXN0Y2xhc3M9IkJlYW5TaGVsbFByZVByb2Nlc3NvciIgdGVzdG5hbWU9IlNldCBEeW5hdHJhY2UgSGVhZGVycyIgZW5hYmxlZD0idHJ1ZSI+CiAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJmaWxlbmFtZSI+PC9zdHJpbmdQcm9wPgogICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0icGFyYW1ldGVycyI+bG9hZC5qbXg8L3N0cmluZ1Byb3A+CiAgICAgICAgICA8Ym9vbFByb3AgbmFtZT0icmVzZXRJbnRlcnByZXRlciI+ZmFsc2U8L2Jvb2xQcm9wPgogICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0ic2NyaXB0Ij4KCmltcG9ydCBvcmcuYXBhY2hlLmptZXRlci51dGlsLkpNZXRlclV0aWxzOwppbXBvcnQgb3JnLmFwYWNoZS5qbWV0ZXIucHJvdG9jb2wuaHR0cC5jb250cm9sLkhlYWRlck1hbmFnZXI7CmltcG9ydCBqYXZhLmlvOwppbXBvcnQgamF2YS51dGlsOwoKLy8gLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQovLyBHZW5lcmF0ZSB0aGUgeC1keW5hdHJhY2UtdGVzdCBoZWFkZXIKLy8gLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLQpTdHJpbmcgTFROPUpNZXRlclV0aWxzLmdldFByb3BlcnR5KCZxdW90O0RUX0xUTiZxdW90Oyk7CmlmKChMVE4gPT0gbnVsbCkgfHwgKExUTi5sZW5ndGgoKSA9PSAwKSkgewogICAgaWYodmFycyAhPSBudWxsKSB7CiAgICAgICAgTFROID0gdmFycy5nZXQoJnF1b3Q7RFRfTFROJnF1b3Q7KTsKICAgIH0KfQppZihMVE4gPT0gbnVsbCkgTFROID0gJnF1b3Q7Tm9UZXN0TmFtZSZxdW90OzsKClN0cmluZyBMU04gPSAoYnNoLmFyZ3MubGVuZ3RoICZndDsgMCkgPyBic2guYXJnc1swXSA6ICZxdW90O1Rlc3QgU2NlbmFyaW8mcXVvdDs7ClN0cmluZyBUU04gPSBzYW1wbGVyLmdldE5hbWUoKTsKU3RyaW5nIFZVID0gY3R4LmdldFRocmVhZEdyb3VwKCkuZ2V0TmFtZSgpICsgY3R4LmdldFRocmVhZE51bSgpOwpTdHJpbmcgaGVhZGVyVmFsdWUgPSAmcXVvdDtMU049JnF1b3Q7KyBMU04gKyAmcXVvdDs7VFNOPSZxdW90OyArIFRTTiArICZxdW90OztMVE49JnF1b3Q7ICsgTFROICsgJnF1b3Q7O1ZVPSZxdW90OyArIFZVICsgJnF1b3Q7OyZxdW90OzsKCi8vIC0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0KLy8gU2V0IGhlYWRlcgovLyAtLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tCkhlYWRlck1hbmFnZXIgaG0gPSBzYW1wbGVyLmdldEhlYWRlck1hbmFnZXIoKTsKaG0ucmVtb3ZlSGVhZGVyTmFtZWQoJnF1b3Q7eC1keW5hdHJhY2UtdGVzdCZxdW90Oyk7CmhtLmFkZChuZXcgb3JnLmFwYWNoZS5qbWV0ZXIucHJvdG9jb2wuaHR0cC5jb250cm9sLkhlYWRlcigmcXVvdDt4LWR5bmF0cmFjZS10ZXN0JnF1b3Q7LCBoZWFkZXJWYWx1ZSkpOwoKICAgICAgICAgIDwvc3RyaW5nUHJvcD4KICAgICAgICA8L0JlYW5TaGVsbFByZVByb2Nlc3Nvcj4KICAgICAgICA8aGFzaFRyZWUvPgogICAgICAgIDxIVFRQU2FtcGxlclByb3h5IGd1aWNsYXNzPSJIdHRwVGVzdFNhbXBsZUd1aSIgdGVzdGNsYXNzPSJIVFRQU2FtcGxlclByb3h5IiB0ZXN0bmFtZT0iaG9tZXBhZ2UiIGVuYWJsZWQ9InRydWUiPgogICAgICAgICAgPGVsZW1lbnRQcm9wIG5hbWU9IkhUVFBzYW1wbGVyLkFyZ3VtZW50cyIgZWxlbWVudFR5cGU9IkFyZ3VtZW50cyIgZ3VpY2xhc3M9IkhUVFBBcmd1bWVudHNQYW5lbCIgdGVzdGNsYXNzPSJBcmd1bWVudHMiIGVuYWJsZWQ9InRydWUiPgogICAgICAgICAgICA8Y29sbGVjdGlvblByb3AgbmFtZT0iQXJndW1lbnRzLmFyZ3VtZW50cyIvPgogICAgICAgICAgPC9lbGVtZW50UHJvcD4KICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkhUVFBTYW1wbGVyLmRvbWFpbiI+JHtfX1AoU0VSVkVSX1VSTCwke1NFUlZFUl9VUkx9KX08L3N0cmluZ1Byb3A+CiAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJIVFRQU2FtcGxlci5wb3J0Ij4ke19fUChTRVJWRVJfUE9SVCwke1NFUlZFUl9QT1JUfSl9PC9zdHJpbmdQcm9wPgogICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iSFRUUFNhbXBsZXIucHJvdG9jb2wiPiR7X19QKFBST1RPQ09MLCR7UFJPVE9DT0x9KX08L3N0cmluZ1Byb3A+CiAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJIVFRQU2FtcGxlci5jb250ZW50RW5jb2RpbmciPjwvc3RyaW5nUHJvcD4KICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkhUVFBTYW1wbGVyLnBhdGgiPi88L3N0cmluZ1Byb3A+CiAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJIVFRQU2FtcGxlci5tZXRob2QiPkdFVDwvc3RyaW5nUHJvcD4KICAgICAgICAgIDxib29sUHJvcCBuYW1lPSJIVFRQU2FtcGxlci5mb2xsb3dfcmVkaXJlY3RzIj50cnVlPC9ib29sUHJvcD4KICAgICAgICAgIDxib29sUHJvcCBuYW1lPSJIVFRQU2FtcGxlci5hdXRvX3JlZGlyZWN0cyI+ZmFsc2U8L2Jvb2xQcm9wPgogICAgICAgICAgPGJvb2xQcm9wIG5hbWU9IkhUVFBTYW1wbGVyLnVzZV9rZWVwYWxpdmUiPnRydWU8L2Jvb2xQcm9wPgogICAgICAgICAgPGJvb2xQcm9wIG5hbWU9IkhUVFBTYW1wbGVyLkRPX01VTFRJUEFSVF9QT1NUIj5mYWxzZTwvYm9vbFByb3A+CiAgICAgICAgICA8Ym9vbFByb3AgbmFtZT0iSFRUUFNhbXBsZXIuQlJPV1NFUl9DT01QQVRJQkxFX01VTFRJUEFSVCI+dHJ1ZTwvYm9vbFByb3A+CiAgICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJIVFRQU2FtcGxlci5lbWJlZGRlZF91cmxfcmUiPjwvc3RyaW5nUHJvcD4KICAgICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9IkhUVFBTYW1wbGVyLmNvbm5lY3RfdGltZW91dCI+PC9zdHJpbmdQcm9wPgogICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iSFRUUFNhbXBsZXIucmVzcG9uc2VfdGltZW91dCI+PC9zdHJpbmdQcm9wPgogICAgICAgIDwvSFRUUFNhbXBsZXJQcm94eT4KICAgICAgICA8aGFzaFRyZWUvPgogICAgICAgIDxDb25zdGFudFRpbWVyIGd1aWNsYXNzPSJDb25zdGFudFRpbWVyR3VpIiB0ZXN0Y2xhc3M9IkNvbnN0YW50VGltZXIiIHRlc3RuYW1lPSJEZWZhdWx0IFRoaW5rIFRpbWUiIGVuYWJsZWQ9InRydWUiPgogICAgICAgICAgPHN0cmluZ1Byb3AgbmFtZT0iQ29uc3RhbnRUaW1lci5kZWxheSI+e19fUChUaGlua1RpbWUsJHtEZWZhdWx0VGhpbmtUaW1lfSl9PC9zdHJpbmdQcm9wPgogICAgICAgIDwvQ29uc3RhbnRUaW1lcj4KICAgICAgICA8aGFzaFRyZWUvPgogICAgICA8L2hhc2hUcmVlPgogICAgICA8UmVzdWx0Q29sbGVjdG9yIGd1aWNsYXNzPSJWaWV3UmVzdWx0c0Z1bGxWaXN1YWxpemVyIiB0ZXN0Y2xhc3M9IlJlc3VsdENvbGxlY3RvciIgdGVzdG5hbWU9IlZpZXcgUmVzdWx0cyBUcmVlIiBlbmFibGVkPSJ0cnVlIj4KICAgICAgICA8Ym9vbFByb3AgbmFtZT0iUmVzdWx0Q29sbGVjdG9yLmVycm9yX2xvZ2dpbmciPmZhbHNlPC9ib29sUHJvcD4KICAgICAgICA8b2JqUHJvcD4KICAgICAgICAgIDxuYW1lPnNhdmVDb25maWc8L25hbWU+CiAgICAgICAgICA8dmFsdWUgY2xhc3M9IlNhbXBsZVNhdmVDb25maWd1cmF0aW9uIj4KICAgICAgICAgICAgPHRpbWU+dHJ1ZTwvdGltZT4KICAgICAgICAgICAgPGxhdGVuY3k+dHJ1ZTwvbGF0ZW5jeT4KICAgICAgICAgICAgPHRpbWVzdGFtcD50cnVlPC90aW1lc3RhbXA+CiAgICAgICAgICAgIDxzdWNjZXNzPnRydWU8L3N1Y2Nlc3M+CiAgICAgICAgICAgIDxsYWJlbD50cnVlPC9sYWJlbD4KICAgICAgICAgICAgPGNvZGU+dHJ1ZTwvY29kZT4KICAgICAgICAgICAgPG1lc3NhZ2U+dHJ1ZTwvbWVzc2FnZT4KICAgICAgICAgICAgPHRocmVhZE5hbWU+dHJ1ZTwvdGhyZWFkTmFtZT4KICAgICAgICAgICAgPGRhdGFUeXBlPnRydWU8L2RhdGFUeXBlPgogICAgICAgICAgICA8ZW5jb2Rpbmc+ZmFsc2U8L2VuY29kaW5nPgogICAgICAgICAgICA8YXNzZXJ0aW9ucz50cnVlPC9hc3NlcnRpb25zPgogICAgICAgICAgICA8c3VicmVzdWx0cz50cnVlPC9zdWJyZXN1bHRzPgogICAgICAgICAgICA8cmVzcG9uc2VEYXRhPmZhbHNlPC9yZXNwb25zZURhdGE+CiAgICAgICAgICAgIDxzYW1wbGVyRGF0YT5mYWxzZTwvc2FtcGxlckRhdGE+CiAgICAgICAgICAgIDx4bWw+ZmFsc2U8L3htbD4KICAgICAgICAgICAgPGZpZWxkTmFtZXM+ZmFsc2U8L2ZpZWxkTmFtZXM+CiAgICAgICAgICAgIDxyZXNwb25zZUhlYWRlcnM+ZmFsc2U8L3Jlc3BvbnNlSGVhZGVycz4KICAgICAgICAgICAgPHJlcXVlc3RIZWFkZXJzPmZhbHNlPC9yZXF1ZXN0SGVhZGVycz4KICAgICAgICAgICAgPHJlc3BvbnNlRGF0YU9uRXJyb3I+ZmFsc2U8L3Jlc3BvbnNlRGF0YU9uRXJyb3I+CiAgICAgICAgICAgIDxzYXZlQXNzZXJ0aW9uUmVzdWx0c0ZhaWx1cmVNZXNzYWdlPmZhbHNlPC9zYXZlQXNzZXJ0aW9uUmVzdWx0c0ZhaWx1cmVNZXNzYWdlPgogICAgICAgICAgICA8YXNzZXJ0aW9uc1Jlc3VsdHNUb1NhdmU+MDwvYXNzZXJ0aW9uc1Jlc3VsdHNUb1NhdmU+CiAgICAgICAgICAgIDxieXRlcz50cnVlPC9ieXRlcz4KICAgICAgICAgICAgPHRocmVhZENvdW50cz50cnVlPC90aHJlYWRDb3VudHM+CiAgICAgICAgICA8L3ZhbHVlPgogICAgICAgIDwvb2JqUHJvcD4KICAgICAgICA8c3RyaW5nUHJvcCBuYW1lPSJmaWxlbmFtZSI+PC9zdHJpbmdQcm9wPgogICAgICA8L1Jlc3VsdENvbGxlY3Rvcj4KICAgICAgPGhhc2hUcmVlLz4KICAgICAgPFJlc3VsdENvbGxlY3RvciBndWljbGFzcz0iU3VtbWFyeVJlcG9ydCIgdGVzdGNsYXNzPSJSZXN1bHRDb2xsZWN0b3IiIHRlc3RuYW1lPSJTdW1tYXJ5IFJlcG9ydCIgZW5hYmxlZD0idHJ1ZSI+CiAgICAgICAgPGJvb2xQcm9wIG5hbWU9IlJlc3VsdENvbGxlY3Rvci5lcnJvcl9sb2dnaW5nIj5mYWxzZTwvYm9vbFByb3A+CiAgICAgICAgPG9ialByb3A+CiAgICAgICAgICA8bmFtZT5zYXZlQ29uZmlnPC9uYW1lPgogICAgICAgICAgPHZhbHVlIGNsYXNzPSJTYW1wbGVTYXZlQ29uZmlndXJhdGlvbiI+CiAgICAgICAgICAgIDx0aW1lPnRydWU8L3RpbWU+CiAgICAgICAgICAgIDxsYXRlbmN5PnRydWU8L2xhdGVuY3k+CiAgICAgICAgICAgIDx0aW1lc3RhbXA+dHJ1ZTwvdGltZXN0YW1wPgogICAgICAgICAgICA8c3VjY2Vzcz50cnVlPC9zdWNjZXNzPgogICAgICAgICAgICA8bGFiZWw+dHJ1ZTwvbGFiZWw+CiAgICAgICAgICAgIDxjb2RlPnRydWU8L2NvZGU+CiAgICAgICAgICAgIDxtZXNzYWdlPnRydWU8L21lc3NhZ2U+CiAgICAgICAgICAgIDx0aHJlYWROYW1lPnRydWU8L3RocmVhZE5hbWU+CiAgICAgICAgICAgIDxkYXRhVHlwZT50cnVlPC9kYXRhVHlwZT4KICAgICAgICAgICAgPGVuY29kaW5nPmZhbHNlPC9lbmNvZGluZz4KICAgICAgICAgICAgPGFzc2VydGlvbnM+dHJ1ZTwvYXNzZXJ0aW9ucz4KICAgICAgICAgICAgPHN1YnJlc3VsdHM+dHJ1ZTwvc3VicmVzdWx0cz4KICAgICAgICAgICAgPHJlc3BvbnNlRGF0YT5mYWxzZTwvcmVzcG9uc2VEYXRhPgogICAgICAgICAgICA8c2FtcGxlckRhdGE+ZmFsc2U8L3NhbXBsZXJEYXRhPgogICAgICAgICAgICA8eG1sPmZhbHNlPC94bWw+CiAgICAgICAgICAgIDxmaWVsZE5hbWVzPnRydWU8L2ZpZWxkTmFtZXM+CiAgICAgICAgICAgIDxyZXNwb25zZUhlYWRlcnM+ZmFsc2U8L3Jlc3BvbnNlSGVhZGVycz4KICAgICAgICAgICAgPHJlcXVlc3RIZWFkZXJzPmZhbHNlPC9yZXF1ZXN0SGVhZGVycz4KICAgICAgICAgICAgPHJlc3BvbnNlRGF0YU9uRXJyb3I+ZmFsc2U8L3Jlc3BvbnNlRGF0YU9uRXJyb3I+CiAgICAgICAgICAgIDxzYXZlQXNzZXJ0aW9uUmVzdWx0c0ZhaWx1cmVNZXNzYWdlPnRydWU8L3NhdmVBc3NlcnRpb25SZXN1bHRzRmFpbHVyZU1lc3NhZ2U+CiAgICAgICAgICAgIDxhc3NlcnRpb25zUmVzdWx0c1RvU2F2ZT4wPC9hc3NlcnRpb25zUmVzdWx0c1RvU2F2ZT4KICAgICAgICAgICAgPGJ5dGVzPnRydWU8L2J5dGVzPgogICAgICAgICAgICA8c2VudEJ5dGVzPnRydWU8L3NlbnRCeXRlcz4KICAgICAgICAgICAgPHVybD50cnVlPC91cmw+CiAgICAgICAgICAgIDx0aHJlYWRDb3VudHM+dHJ1ZTwvdGhyZWFkQ291bnRzPgogICAgICAgICAgICA8aWRsZVRpbWU+dHJ1ZTwvaWRsZVRpbWU+CiAgICAgICAgICAgIDxjb25uZWN0VGltZT50cnVlPC9jb25uZWN0VGltZT4KICAgICAgICAgIDwvdmFsdWU+CiAgICAgICAgPC9vYmpQcm9wPgogICAgICAgIDxzdHJpbmdQcm9wIG5hbWU9ImZpbGVuYW1lIj48L3N0cmluZ1Byb3A+CiAgICAgIDwvUmVzdWx0Q29sbGVjdG9yPgogICAgICA8aGFzaFRyZWUvPgogICAgPC9oYXNoVHJlZT4KICA8L2hhc2hUcmVlPgo8L2ptZXRlclRlc3RQbGFuPgo="
	resourceUri := "load.jmx"
	resourceUriPath := "/"+resourceUri

	createProjectRequest := models.Project{
		ProjectName:  projectName,
	}

	createStageRequests := []models.Stage{
		{
			StageName: "dev",
		},
		{
			StageName: "production",
		},
	}

	createServiceRequests := []models.Service{
		{
			ServiceName: "app",
		},
		{
			ServiceName: "app-db",
		},
	}

	createResourceRequest := models.Resources{
		Resources: []*models.Resource{
			{
				ResourceContent: resourceContent,
				ResourceURI:     &resourceUri,
			},
		},
	}

	updateProjectRequest := models.Project{
		ProjectName:  projectName,
		GitUser: "some_random_git_user",
	}

	updateNonExistingProjectRequest := models.Project{
		ProjectName:  nonExistingProjectName,
		GitUser: "some_random_git_user",
	}

	///////////////////////////////////////
	// Creation of objects
	///////////////////////////////////////

	t.Logf("Creating a new project %s", projectName)
	resp, err := ApiPOSTRequest("/configuration-service/v1/project", createProjectRequest, 3)
	require.Nil(t, err)
	require.Equal(t, 204, resp.Response().StatusCode)

	t.Logf("Creating a new resource for project %s", projectName)
	resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+projectName+"/resource", createResourceRequest, 3)
	require.Nil(t, err)
	require.Equal(t, 201, resp.Response().StatusCode)

	t.Logf("Checking resource for project %s", projectName)
	resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/resource"+resourceUriPath, 3)
	require.Nil(t, err)
	require.Equal(t, 200, resp.Response().StatusCode)

	t.Logf("Checking body of the received response")
	resource := models.Resource{}
	err = resp.ToJSON(&resource)
	require.Nil(t, err)
	require.Equal(t, resourceUri, *resource.ResourceURI)

	t.Logf("Checking all resources for project %s", projectName)
	resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/resource", 3)
	require.Nil(t, err)
	require.Equal(t, 200, resp.Response().StatusCode)

	t.Logf("Checking body of the received response")
	resources := models.Resources{}
	err = resp.ToJSON(&resources)
	require.Nil(t, err)
	require.Equal(t, float64(2), resources.TotalCount)
	require.Equal(t, resourceUriPath, *resources.Resources[0].ResourceURI)

	t.Logf("Checking all resources for non-existing project %s", nonExistingProjectName)
	resp, err = ApiGETRequest("/configuration-service/v1/project/"+nonExistingProjectName+"/resource", 3)
	require.Nil(t, err)
	require.Equal(t, 404, resp.Response().StatusCode)

	t.Logf("Creating an existing new project %s", projectName)
	resp, err = ApiPOSTRequest("/configuration-service/v1/project", createProjectRequest, 3)
	require.Nil(t, err)
	require.Equal(t, 400, resp.Response().StatusCode)

	t.Logf("Creating a new resource for non-existing project %s", nonExistingProjectName)
	resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+nonExistingProjectName+"/resource", createResourceRequest, 3)
	require.Nil(t, err)
	require.Equal(t, 400, resp.Response().StatusCode)

	for _, stageReq := range createStageRequests {
		t.Logf("Creating a new stage %s in project %s", stageReq.StageName, projectName)
		resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+projectName+"/stage", stageReq, 3)
		require.Nil(t, err)
		require.Equal(t, 204, resp.Response().StatusCode)

		t.Logf("Creating a new resource for stage %s for project %s", stageReq.StageName, projectName)
		resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/resource", createResourceRequest, 3)
		require.Nil(t, err)
		require.Equal(t, 201, resp.Response().StatusCode)

		t.Logf("Checking resource for stage %s for project %s", stageReq.StageName, projectName)
		resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/resource"+resourceUriPath, 3)
		require.Nil(t, err)
		require.Equal(t, 200, resp.Response().StatusCode)

		t.Logf("Checking body of the received response")
		resource := models.Resource{}
		err = resp.ToJSON(&resource)
		require.Nil(t, err)
		require.Equal(t, resourceUri, *resource.ResourceURI)

		t.Logf("Checking all resources for stage %s for project %s", stageReq.StageName, projectName)
		resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/resource", 3)
		require.Nil(t, err)
		require.Equal(t, 200, resp.Response().StatusCode)

		t.Logf("Checking body of the received response")
		resources := models.Resources{}
		err = resp.ToJSON(&resources)
		require.Nil(t, err)
		require.Equal(t, float64(2), resources.TotalCount)
		require.Equal(t, resourceUriPath, *resources.Resources[0].ResourceURI)

		t.Logf("Checking all resources for non-existing stage %s for project %s", nonExistingStageName, projectName)
		resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+nonExistingStageName+"/resource", 3)
		require.Nil(t, err)
		require.Equal(t, 404, resp.Response().StatusCode)

		t.Logf("Creating an existing new stage %s in project %s", stageReq.StageName, projectName)
		resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+projectName+"/stage", stageReq, 3)
		require.Nil(t, err)
		require.Equal(t, 400, resp.Response().StatusCode)

		t.Logf("Creating a new resource for non-existing stage %s for project %s", nonExistingStageName, projectName)
		resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+projectName+"/stage/"+nonExistingStageName+"/resource", createResourceRequest, 3)
		require.Nil(t, err)
		require.Equal(t, 400, resp.Response().StatusCode)
	}

	for _, stageReq := range createStageRequests {
		for _, serviceReq := range createServiceRequests {
			t.Logf("Creating a new service %s in stage %s in project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service", serviceReq, 3)
			require.Nil(t, err)
			require.Equal(t, 204, resp.Response().StatusCode)

			t.Logf("Creating a new resource for service %s in stage %s for project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+serviceReq.ServiceName+"/resource", createResourceRequest, 3)
			require.Nil(t, err)
			require.Equal(t, 201, resp.Response().StatusCode)

			t.Logf("Checking resource for service %s in stage %s for project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+serviceReq.ServiceName+"/resource"+resourceUriPath, 3)
			require.Nil(t, err)
			require.Equal(t, 200, resp.Response().StatusCode)

			t.Logf("Checking body of the received response")
			resource := models.Resource{}
			err = resp.ToJSON(&resource)
			require.Nil(t, err)
			require.Equal(t, resourceUri, *resource.ResourceURI)

			t.Logf("Checking all resources for service %s in stage %s for project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+serviceReq.ServiceName+"/resource", 3)
			require.Nil(t, err)
			require.Equal(t, 200, resp.Response().StatusCode)

			t.Logf("Checking body of the received response")
			resources := models.Resources{}
			err = resp.ToJSON(&resources)
			require.Nil(t, err)
			require.Equal(t, float64(2), resources.TotalCount)
			require.Equal(t, resourceUriPath, *resources.Resources[0].ResourceURI)

			t.Logf("Checking all resources for non-existing service %s in stage %s for project %s", nonExistingServiceName, stageReq.StageName, projectName)
			resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+nonExistingServiceName+"/resource", 3)
			require.Nil(t, err)
			require.Equal(t, 404, resp.Response().StatusCode)

			t.Logf("Creating an existing new service %s in stage %s in project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service", serviceReq, 3)
			require.Nil(t, err)
			require.Equal(t, 400, resp.Response().StatusCode)

			t.Logf("Creating a new resource for non-existing service %s in stage %s for project %s", nonExistingServiceName, stageReq.StageName, projectName)
			resp, err = ApiPOSTRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+nonExistingServiceName+"/resource", createResourceRequest, 3)
			require.Nil(t, err)
			require.Equal(t, 400, resp.Response().StatusCode)
		}
	}

	///////////////////////////////////////
	// Update of objects
	///////////////////////////////////////

	t.Logf("Updating project %s", projectName)
	resp, err = ApiPUTRequest("/configuration-service/v1/project/"+projectName, updateProjectRequest, 3)
	require.Nil(t, err)
	require.Equal(t, 204, resp.Response().StatusCode)

	t.Logf("Updating non existing project %s", nonExistingProjectName)
	resp, err = ApiPUTRequest("/configuration-service/v1/project/"+nonExistingProjectName, updateNonExistingProjectRequest, 3)
	require.Nil(t, err)
	require.Equal(t, 400, resp.Response().StatusCode)

	for _, stageReq := range createStageRequests {
		t.Logf("Updating stage %s in project %s", stageReq.StageName, projectName)
		resp, err = ApiPUTRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName, stageReq, 3)
		require.Nil(t, err)
		require.Equal(t, 501, resp.Response().StatusCode) // should be 204 in resource-service

	}

	for _, stageReq := range createStageRequests {
		for _, serviceReq := range createServiceRequests {
			t.Logf("Updating service %s in stage %s in project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiPUTRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+serviceReq.ServiceName, serviceReq, 3)
			require.Nil(t, err)
			require.Equal(t, 501, resp.Response().StatusCode) // should be 204 in resource-service

		}
	}

	///////////////////////////////////////
	// Deletion of objects
	///////////////////////////////////////

	for _, stageReq := range createStageRequests {
		for _, serviceReq := range createServiceRequests {
			t.Logf("Deleting the resource from service %s from stage %s from project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+serviceReq.ServiceName+"/resource"+resourceUriPath, 3)
			require.Nil(t, err)
			require.Equal(t, 204, resp.Response().StatusCode)

			t.Logf("Checking non-existing resource for service %s for stage %s for project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+serviceReq.ServiceName+"/resource"+resourceUriPath, 3)
			require.Nil(t, err)
			require.Equal(t, 404, resp.Response().StatusCode)

			t.Logf("Deleting non-existing resource from service %s from stage %s from project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+serviceReq.ServiceName+"/resource"+resourceUriPath, 3)
			require.Nil(t, err)
			require.Equal(t, 500, resp.Response().StatusCode) //chce to iny kod

			t.Logf("Deleting service %s in stage %s in project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+serviceReq.ServiceName, 3)
			require.Nil(t, err)
			require.Equal(t, 204, resp.Response().StatusCode)

			t.Logf("Checking resource for non-existing service %s in stage %s for project %s", serviceReq.ServiceName, stageReq.StageName, projectName)
			resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+serviceReq.ServiceName+"/resource"+resourceUriPath, 3)
			require.Nil(t, err)
			require.Equal(t, 404, resp.Response().StatusCode)

			t.Logf("Deleting non-existing service %s in stage %s in project %s", nonExistingServiceName, stageReq.StageName, projectName)
			resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/service/"+nonExistingServiceName, 3)
			require.Nil(t, err)
			require.Equal(t, 400, resp.Response().StatusCode)
		}
	}

	for _, stageReq := range createStageRequests {
		t.Logf("Deleting the resource from stage %s from project %s", stageReq.StageName, projectName)
		resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/resource"+resourceUriPath, 3)
		require.Nil(t, err)
		require.Equal(t, 204, resp.Response().StatusCode)

		t.Logf("Checking non-existing resource for stage %s for project %s", stageReq.StageName, projectName)
		resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/resource"+resourceUriPath, 3)
		require.Nil(t, err)
		require.Equal(t, 404, resp.Response().StatusCode)

		t.Logf("Deleting non-existing resource from stage %s from project %s", stageReq.StageName, projectName)
		resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/resource"+resourceUriPath, 3)
		require.Nil(t, err)
		require.Equal(t, 500, resp.Response().StatusCode) //chce to iny kod

		t.Logf("Deleting stage %s in project %s", stageReq.StageName, projectName)
		resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName, 3)
		require.Nil(t, err)
		require.Equal(t, 501, resp.Response().StatusCode) //will be 204 for restore-service

		t.Logf("Checking resource for non-existing stage %s for project %s", stageReq.StageName, projectName)
		resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/stage/"+stageReq.StageName+"/resource"+resourceUriPath, 3)
		require.Nil(t, err)
		require.Equal(t, 404, resp.Response().StatusCode)

		//delete non-existing project
		t.Logf("Deleting non-existing stage %s in project %s", nonExistingStageName, projectName)
		resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/stage/"+nonExistingStageName, 3)
		require.Nil(t, err)
		require.Equal(t, 501, resp.Response().StatusCode) //will be 400 for restore-service
	}

	t.Logf("Deleting the resource from project %s", projectName)
	resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/resource"+resourceUriPath, 3)
	require.Nil(t, err)
	require.Equal(t, 204, resp.Response().StatusCode)

	t.Logf("Checking non-existing resource for project %s", projectName)
	resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/resource"+resourceUriPath, 3)
	require.Nil(t, err)
	require.Equal(t, 404, resp.Response().StatusCode)

	t.Logf("Deleting non-existing resource from project %s", projectName)
	resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName+"/resource"+resourceUriPath, 3)
	require.Nil(t, err)
	require.Equal(t, 500, resp.Response().StatusCode) //chce to iny kod

	t.Logf("Deleting the project %s", projectName)
	resp, err = ApiDELETERequest("/configuration-service/v1/project/"+projectName, 3)
	require.Nil(t, err)
	require.Equal(t, 204, resp.Response().StatusCode)

	t.Logf("Checking resource for non-existing project %s", projectName)
	resp, err = ApiGETRequest("/configuration-service/v1/project/"+projectName+"/resource"+resourceUriPath, 3)
	require.Nil(t, err)
	require.Equal(t, 404, resp.Response().StatusCode)

	t.Logf("Deleting non-existing project %s", nonExistingProjectName)
	resp, err = ApiDELETERequest("/configuration-service/v1/project/"+nonExistingProjectName, 3)
	require.Nil(t, err)
	require.Equal(t, 204, resp.Response().StatusCode)

}
