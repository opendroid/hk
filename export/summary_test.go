package export

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)
// TestHealthData_RecordsSummary
// go test -run TestHealthData_RecordsSummary -v
func TestHealthData_RecordsSummary(t *testing.T) {
	file := "../public/sampleexport.xml"
	d, err := Parse(file)
	if err != nil {
		t.Fail()
		return
	}
	if d == nil {
		t.Fail()
		return
	}
	d.DescribeRecords()
	j, err := json.Marshal(d)
	want := `{"locate":"en_US","exported_date":{"value":"2020-12-06 20:24:55 -0800"},"me":{"dob":"1990-12-06","biological_sex":"HKBiologicalSexMale","blood_type":"HKBloodTypeNotSet","skin_type":"HKFitzpatrickSkinTypeNotSet","medication_use":"None"},"records":[{"type":"HKQuantityTypeIdentifierBodyMassIndex","unit":"count","value":"22.4","source_name":"Renpho","source_version":"4","creation_date":"2020-04-19 06:53:41 -0800","start_date":"2020-04-18 05:52:26 -0800","end_date":"2020-04-18 05:52:26 -0800"},{"type":"HKQuantityTypeIdentifierBodyMassIndex","unit":"count","value":"22.6","source_name":"Renpho","source_version":"4","creation_date":"2020-04-19 06:53:41 -0800","start_date":"2020-04-17 05:57:55 -0800","end_date":"2020-04-17 05:57:55 -0800"},{"type":"HKQuantityTypeIdentifierBodyMassIndex","unit":"count","value":"22.4","source_name":"Renpho","source_version":"4","creation_date":"2020-04-19 06:53:44 -0800","start_date":"2020-04-15 05:58:55 -0800","end_date":"2020-04-15 05:58:55 -0800"},{"type":"HKQuantityTypeIdentifierBodyMassIndex","unit":"count","value":"22.2","source_name":"Renpho","source_version":"4","creation_date":"2020-04-19 06:53:45 -0800","start_date":"2020-04-14 14:48:12 -0800","end_date":"2020-04-14 14:48:12 -0800"},{"type":"HKQuantityTypeIdentifierHeartRateVariabilitySDNN","unit":"ms","value":"26.467","source_name":"🕛🕐🕑🕒","source_version":"7.1","device":"\u003c\u003cHKDevice: 0x28347b340\u003e, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch4,2, software:7.1\u003e","creation_date":"2020-12-06 11:27:05 -0800","start_date":"2020-12-06 11:26:00 -0800","end_date":"2020-12-06 11:27:05 -0800","hrv":[{"instantaneous_bpm":[{"bpm":"58","time":"11:26:02.34 AM"},{"bpm":"58","time":"11:26:03.37 AM"},{"bpm":"58","time":"11:26:04.41 AM"},{"bpm":"59","time":"11:26:05.42 AM"},{"bpm":"59","time":"11:26:06.44 AM"},{"bpm":"60","time":"11:26:07.44 AM"},{"bpm":"59","time":"11:26:08.46 AM"},{"bpm":"58","time":"11:26:09.49 AM"},{"bpm":"61","time":"11:26:10.48 AM"},{"bpm":"57","time":"11:26:11.52 AM"}]}]}],"clinical_records":[{"type":"Procedure","id":"TkxvzAawHeQb3Fi8PI60qLcpVWSuan8Gk1dSZsdZkm.YB","source_name":"Sutter Health","source_url":"https://apiservices.sutterhealth.org/ifs/api/FHIR/DSTU2/Procedure/TkxvzAawHeQb3Fi8PI60qLcpVWSuan8Gk1dSZsdZkm.YB","fhir_version":"1.0.2","received_date":"2020-12-06 20:24:28 -0800","file_path":"/clinical-records/Procedure-488D59AE-81D4-4597-8FBB-713DE1CEF228.json"},{"type":"Procedure","id":"Tu.9MGYy8koq4Z8.LtMbhLYYsDN0RNAhf1mAYIP3B5ewB","source_name":"Sutter Health","source_url":"https://apiservices.sutterhealth.org/ifs/api/FHIR/DSTU2/Procedure/Tu.9MGYy8koq4Z8.LtMbhLYYsDN0RNAhf1mAYIP3B5ewB","fhir_version":"1.0.2","received_date":"2020-12-06 20:24:28 -0800","file_path":"/clinical-records/Procedure-B4F159B1-E990-49D5-AEAD-CB6C01F99D08.json"},{"type":"Procedure","id":"T2WHl11o4AsRpxTDldMctQkrGO7gKJJRgpshqYFrMngsB","source_name":"Sutter Health","source_url":"https://apiservices.sutterhealth.org/ifs/api/FHIR/DSTU2/Procedure/T2WHl11o4AsRpxTDldMctQkrGO7gKJJRgpshqYFrMngsB","fhir_version":"1.0.2","received_date":"2020-12-06 20:24:28 -0800","file_path":"/clinical-records/Procedure-750A37FB-42CB-463A-BAAD-8730AE96FD1D.json"},{"type":"Immunization","id":"Tdot1l.AKI5tXy3lOznArzgB","source_name":"Sutter Health","source_url":"https://apiservices.sutterhealth.org/ifs/api/FHIR/DSTU2/Immunization/Tdot1l.AKI5tXy3lOznArzgB","fhir_version":"1.0.2","received_date":"2020-12-06 20:24:29 -0800","file_path":"/clinical-records/Immunization-FD9A852A-4601-494A-B97F-F5CF7A4082E4.json"},{"type":"DiagnosticReport","id":"T18EOBCL62BP9Ym5ddQzK9XS4fVSCMjR2n6K-7k8eOesB","source_name":"Sutter Health","source_url":"https://apiservices.sutterhealth.org/ifs/api/FHIR/DSTU2/DiagnosticReport/T18EOBCL62BP9Ym5ddQzK9XS4fVSCMjR2n6K-7k8eOesB","fhir_version":"1.0.2","received_date":"2020-08-06 05:01:22 -0800","file_path":"/clinical-records/DiagnosticReport-F011639A-7CDB-4E7F-B195-E4DFF6D2FB76.json"},{"type":"Patient","id":"T.2KQWgolUTU7Mj4P37ABZ5DQ.8nOC4KyM5szJtBRJW4B","source_name":"Sutter Health","source_url":"https://apiservices.sutterhealth.org/ifs/api/FHIR/DSTU2/Patient/T.2KQWgolUTU7Mj4P37ABZ5DQ.8nOC4KyM5szJtBRJW4B","fhir_version":"1.0.2","received_date":"2020-12-06 20:24:30 -0800","file_path":"/clinical-records/Patient-6221D445-19F5-4DBB-A995-84FCAFF6F3B5.json"}],"workouts":[{"activity_type":"HKWorkoutActivityTypeWalking","duration":"53.05183674693107","duration_unit":"min","total_distance":"2.985843180268872","total_distance_unit":"mi","total_energy_burned":"259.191938540603","total_energy_burned_unit":"Cal","source_name":"🕛🕐🕑🕒","source_version":"5.1.2","device":"\u003c\u003cHKDevice: 0x2835f0aa0\u003e, name:Apple Watch, manufacturer:Apple, model:Watch, hardware:Watch4,2, software:5.1.2\u003e","creation_date":"2018-12-17 08:48:03 -0800","start_date":"2018-12-17 07:54:59 -0800","end_date":"2018-12-17 08:48:02 -0800","metadata":[{"key":"HKIndoorWorkout","value":"0"}],"events":[{"type":"HKWorkoutEventTypeSegment","date":"2018-12-17 07:54:59 -0800","duration":"11.80219223300616","duration_unit":"min"},{"type":"HKWorkoutEventTypeSegment","date":"2018-12-17 07:54:59 -0800","duration":"17.86551037430763","duration_unit":"min"},{"type":"HKWorkoutEventTypeSegment","date":"2018-12-17 08:06:47 -0800","duration":"10.07729632655779","duration_unit":"min"},{"type":"HKWorkoutEventTypeSegment","date":"2018-12-17 08:12:51 -0800","duration":"17.20985964139303","duration_unit":"min"},{"type":"HKWorkoutEventTypeSegment","date":"2018-12-17 08:16:51 -0800","duration":"10.97524305184682","duration_unit":"min"},{"type":"HKWorkoutEventTypeSegment","date":"2018-12-17 08:27:50 -0800","duration":"10.93215700984001","duration_unit":"min"},{"type":"HKWorkoutEventTypeSegment","date":"2018-12-17 08:30:03 -0800","duration":"17.89234066009521","duration_unit":"min"},{"type":"HKWorkoutEventTypeSegment","date":"2018-12-17 08:38:46 -0800","duration":"9.180822054545084","duration_unit":"min"}],"routes":[{"source_name":"🕛🕐🕑🕒","source_version":"7.1","creation_date":"2020-12-06 15:03:11 -0800","start_date":"2020-12-06 13:13:52 -0800","end_date":"2020-12-06 15:02:58 -0800","metadata":[{"key":"HKMetadataKeySyncVersion","value":"2"},{"key":"HKMetadataKeySyncIdentifier","value":"242FD852-9322-43FE-8A2D-B6416872C73D"}],"file_reference":[{"path":"/workout-routes/route_2020-12-06_3.02pm.gpx"}]}]}],"activities_summary":[{"yyyymmdd":"2020-11-20","energy_burned":"487.231","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"74","exercise_time_goal":"30","stand_hours":"16","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-21","energy_burned":"551.951","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"79","exercise_time_goal":"30","stand_hours":"16","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-22","energy_burned":"534.026","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"80","exercise_time_goal":"30","stand_hours":"18","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-23","energy_burned":"585.547","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"83","exercise_time_goal":"30","stand_hours":"17","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-24","energy_burned":"604.823","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"83","exercise_time_goal":"30","stand_hours":"19","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-25","energy_burned":"606.726","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"85","exercise_time_goal":"30","stand_hours":"18","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-26","energy_burned":"599.972","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"88","exercise_time_goal":"30","stand_hours":"18","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-27","energy_burned":"648.394","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"100","exercise_time_goal":"30","stand_hours":"18","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-28","energy_burned":"758.974","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"141","exercise_time_goal":"30","stand_hours":"18","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-29","energy_burned":"585.901","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"102","exercise_time_goal":"30","stand_hours":"19","stand_hours_goal":"11"},{"yyyymmdd":"2020-11-30","energy_burned":"695.816","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"109","exercise_time_goal":"30","stand_hours":"19","stand_hours_goal":"11"},{"yyyymmdd":"2020-12-01","energy_burned":"655.492","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"95","exercise_time_goal":"30","stand_hours":"16","stand_hours_goal":"11"},{"yyyymmdd":"2020-12-02","energy_burned":"634.471","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"90","exercise_time_goal":"30","stand_hours":"18","stand_hours_goal":"11"},{"yyyymmdd":"2020-12-03","energy_burned":"659.84","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"96","exercise_time_goal":"30","stand_hours":"17","stand_hours_goal":"11"},{"yyyymmdd":"2020-12-04","energy_burned":"647.78","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"93","exercise_time_goal":"30","stand_hours":"18","stand_hours_goal":"11"},{"yyyymmdd":"2020-12-05","energy_burned":"583.725","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"91","exercise_time_goal":"30","stand_hours":"16","stand_hours_goal":"11"},{"yyyymmdd":"2020-12-06","energy_burned":"610.788","energy_burned_goal":"300","energy_burned_unit":"Cal","move_time":"0","move_time_goal":"0","exercise_time":"113","exercise_time_goal":"30","stand_hours":"14","stand_hours_goal":"11"}]}`
	assert.Nil(t, err)
	assert.Equal(t, want, string(j))

	// Look at summary
	summary := d.RecordsSummary()
	j, err = json.Marshal(summary)
	assert.Nil(t, err)
	assert.NotNil(t, j)
	t.Logf("Summary: %s\n", string(j))
}