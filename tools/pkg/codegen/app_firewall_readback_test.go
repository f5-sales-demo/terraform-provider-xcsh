package codegen

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func wafReadbackAttributes() []openapi.TerraformAttribute {
	marker := func(name, goName string) openapi.TerraformAttribute {
		return openapi.TerraformAttribute{Name: name, GoName: goName, TfsdkTag: name, JsonName: name, Type: "object", EmptyObjectMarker: true, Optional: true, IsSpecField: true}
	}
	block := func(name, goName string, children ...openapi.TerraformAttribute) openapi.TerraformAttribute {
		return openapi.TerraformAttribute{Name: name, GoName: goName, TfsdkTag: name, JsonName: name, IsBlock: true, NestedBlockType: "single", Optional: true, IsSpecField: true, NestedAttributes: children}
	}
	return []openapi.TerraformAttribute{
		marker("disable_ai_enhancements", "DisableAiEnhancements"),
		block("enable_ai_enhancements", "EnableAiEnhancements", marker("mitigate_high_risk_action", "MitigateHighRiskAction"), marker("mitigate_high_medium_risk_action", "MitigateHighMediumRiskAction")),
		block("detection_settings", "DetectionSettings", marker("enable_suppression", "EnableSuppression"), marker("disable_suppression", "DisableSuppression"), marker("enable_threat_campaigns", "EnableThreatCampaigns"), marker("disable_threat_campaigns", "DisableThreatCampaigns"), block("signature_selection_setting", "SignatureSelectionSetting", block("signature_settings_by_accuracy", "SignatureSettingsByAccuracy", openapi.TerraformAttribute{Name: "high_accuracy_action", GoName: "HighAccuracyAction", TfsdkTag: "high_accuracy_action", JsonName: "high_accuracy_action", Type: "string", Optional: true}, openapi.TerraformAttribute{Name: "medium_accuracy_action", GoName: "MediumAccuracyAction", TfsdkTag: "medium_accuracy_action", JsonName: "medium_accuracy_action", Type: "string", Optional: true}, openapi.TerraformAttribute{Name: "low_accuracy_action", GoName: "LowAccuracyAction", TfsdkTag: "low_accuracy_action", JsonName: "low_accuracy_action", Type: "string", Optional: true})), marker("default_violation_settings", "DefaultViolationSettings"), block("violation_settings", "ViolationSettings", openapi.TerraformAttribute{Name: "disabled_violation_types", GoName: "DisabledViolationTypes", TfsdkTag: "disabled_violation_types", JsonName: "disabled_violation_types", Type: "list", ElementType: "string", Optional: true})),
	}
}

// Execute the actual generated decoder, rather than treating an unrelated plan
// update as evidence that the changed AI choice was read faithfully (#2573).
func TestAppFirewallGeneratedReadback(t *testing.T) {
	attrs := wafReadbackAttributes()
	code, err := RenderSpecUnmarshalCode(attrs, "\t", "AppFirewall")
	if err != nil {
		t.Fatal(err)
	}
	root := repoRootFromTest(t)
	dir, err := os.MkdirTemp(filepath.Join(root, "internal"), "zz_wafreadback_")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	source := `package probe
import (
 "context"
 "testing"
 "github.com/f5-sales-demo/terraform-provider-xcsh/internal/normalization"
 "github.com/hashicorp/terraform-plugin-framework/attr"
 "github.com/hashicorp/terraform-plugin-framework/diag"
 "github.com/hashicorp/terraform-plugin-framework/types"
)
` + RenderNestedModelTypes("AppFirewall", attrs) + `
type Model struct {
 DisableAiEnhancements types.Object
 EnableAiEnhancements *AppFirewallEnableAiEnhancementsModel
 DetectionSettings *AppFirewallDetectionSettingsModel
}
func read(data *Model, spec map[string]interface{}, isImport bool) {
 ctx := context.Background()
 resp := struct{Diagnostics diag.Diagnostics}{}
 _,_ = ctx,resp
 apiResource := struct{Spec map[string]interface{}}{spec}
` + code + `
}
func marker() types.Object {return types.ObjectValueMust(map[string]attr.Type{},map[string]attr.Value{})}
func absent() types.Object {return types.ObjectNull(map[string]attr.Type{})}
func high() *Model {return &Model{DisableAiEnhancements:absent(),EnableAiEnhancements:&AppFirewallEnableAiEnhancementsModel{MitigateHighRiskAction:marker(),MitigateHighMediumRiskAction:absent()}}}
func TestAITransitions(t *testing.T) {
 for _, imported := range []bool{false,true} {
  t.Run(map[bool]string{false:"refresh",true:"import"}[imported],func(t *testing.T){
   d:=high()
   read(d,map[string]interface{}{"enable_ai_enhancements":map[string]interface{}{"mitigate_high_medium_risk_action":map[string]interface{}{}}},imported)
   if d.EnableAiEnhancements==nil || !d.EnableAiEnhancements.MitigateHighRiskAction.IsNull() || d.EnableAiEnhancements.MitigateHighMediumRiskAction.IsNull(){t.Fatal("High -> High/Medium was not read")}
   read(d,map[string]interface{}{"disable_ai_enhancements":map[string]interface{}{}},false)
   if d.EnableAiEnhancements!=nil {t.Fatal("enabled -> disabled retained AI block")}
   read(d,map[string]interface{}{"enable_ai_enhancements":map[string]interface{}{"mitigate_high_risk_action":map[string]interface{}{}}},false)
   if d.EnableAiEnhancements==nil || d.EnableAiEnhancements.MitigateHighRiskAction.IsNull() || !d.DisableAiEnhancements.IsNull(){t.Fatal("disabled -> enabled was not read")}
  })
 }
}
func TestExplicitDisabledToEnabled(t *testing.T) {
 d:=&Model{DisableAiEnhancements:marker()}
 read(d,map[string]interface{}{"enable_ai_enhancements":map[string]interface{}{"mitigate_high_risk_action":map[string]interface{}{}}},false)
 if d.EnableAiEnhancements==nil || !d.DisableAiEnhancements.IsNull(){t.Fatal("configured disabled arm concealed live enable")}
 read(d,map[string]interface{}{"disable_ai_enhancements":map[string]interface{}{}},false)
 if d.EnableAiEnhancements!=nil {t.Fatal("enabled arm retained after disable")}
}
func TestEquivalentViolationExpansion(t *testing.T) {
 values:=[]interface{}{
 "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST","VIOL_COOKIE_MALFORMED","VIOL_ENCODING",
 "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS","VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING","VIOL_EVASION_BAD_UNESCAPE",
 }
 expanded:=map[string]interface{}{"detection_settings":map[string]interface{}{"violation_settings":map[string]interface{}{"disabled_violation_types":values}}}
 for _,imported:=range []bool{false,true} {
  d:=&Model{DisableAiEnhancements:absent(),DetectionSettings:&AppFirewallDetectionSettingsModel{DefaultViolationSettings:marker()}}
  for i:=0;i<2;i++ {
   read(d,expanded,imported)
   if d.DetectionSettings.DefaultViolationSettings.IsNull() || d.DetectionSettings.ViolationSettings!=nil {t.Fatal("equivalent default expansion failed convergence")}
  }
 }
 d:=&Model{DisableAiEnhancements:absent(),DetectionSettings:&AppFirewallDetectionSettingsModel{DefaultViolationSettings:absent(),ViolationSettings:&AppFirewallDetectionSettingsViolationSettingsModel{}}}
 read(d,expanded,false)
 if !d.DetectionSettings.DefaultViolationSettings.IsNull() || d.DetectionSettings.ViolationSettings==nil {t.Fatal("explicit list lost its representation")}
}
func TestViolationChoiceDrift(t *testing.T) {
 d:=&Model{DisableAiEnhancements:absent(),DetectionSettings:&AppFirewallDetectionSettingsModel{DefaultViolationSettings:marker()}}
 read(d,map[string]interface{}{"detection_settings":map[string]interface{}{"violation_settings":map[string]interface{}{"disabled_violation_types":[]interface{}{"VIOL_ENCODING"}}}},false)
 if !d.DetectionSettings.DefaultViolationSettings.IsNull() || d.DetectionSettings.ViolationSettings==nil {t.Fatal("real violation drift hidden or conflicting choices retained")}
 read(d,map[string]interface{}{"detection_settings":map[string]interface{}{"default_violation_settings":map[string]interface{}{}}},false)
 if d.DetectionSettings.DefaultViolationSettings.IsNull() || d.DetectionSettings.ViolationSettings!=nil {t.Fatal("violation choice restoration failed")}
}
func defaults() map[string]interface{} {return map[string]interface{}{
 "enable_suppression":map[string]interface{}{}, "enable_threat_campaigns":map[string]interface{}{},
 "signature_selection_setting":map[string]interface{}{"signature_settings_by_accuracy":map[string]interface{}{"high_accuracy_action":"SIG_BLOCK","medium_accuracy_action":"SIG_BLOCK","low_accuracy_action":"SIG_IGNORE"}},
}}
func minimal() *Model {return &Model{DisableAiEnhancements:absent(),DetectionSettings:&AppFirewallDetectionSettingsModel{EnableSuppression:absent(),DisableSuppression:absent(),EnableThreatCampaigns:absent(),DisableThreatCampaigns:absent(),DefaultViolationSettings:absent(),SignatureSelectionSetting:&AppFirewallDetectionSettingsSignatureSelectionSettingModel{}}}}
func TestImplicitDefaultImportRefresh(t *testing.T) {
 for _,imported:=range []bool{false,true} {
  d:=minimal()
  for i:=0;i<2;i++ {
   read(d,map[string]interface{}{"detection_settings":defaults()},imported && i==0)
   if !d.DetectionSettings.EnableSuppression.IsNull() || !d.DetectionSettings.EnableThreatCampaigns.IsNull() || d.DetectionSettings.SignatureSelectionSetting.SignatureSettingsByAccuracy!=nil {t.Fatal("omitted server defaults materialized in state")}
  }
 }
}
func TestManagedDefaultPolicyDrift(t *testing.T) {
 d:=minimal()
 d.DetectionSettings.EnableSuppression=marker()
 d.DetectionSettings.EnableThreatCampaigns=marker()
 d.DetectionSettings.SignatureSelectionSetting.SignatureSettingsByAccuracy=&AppFirewallDetectionSettingsSignatureSelectionSettingSignatureSettingsByAccuracyModel{HighAccuracyAction:types.StringValue("SIG_BLOCK"),MediumAccuracyAction:types.StringValue("SIG_BLOCK"),LowAccuracyAction:types.StringValue("SIG_IGNORE")}
 wire:=defaults()
 delete(wire,"enable_suppression");wire["disable_suppression"]=map[string]interface{}{}
 delete(wire,"enable_threat_campaigns");wire["disable_threat_campaigns"]=map[string]interface{}{}
 wire["signature_selection_setting"].(map[string]interface{})["signature_settings_by_accuracy"].(map[string]interface{})["high_accuracy_action"]="SIG_IGNORE"
 read(d,map[string]interface{}{"detection_settings":wire},false)
 if !d.DetectionSettings.EnableSuppression.IsNull() || d.DetectionSettings.DisableSuppression.IsNull() || !d.DetectionSettings.EnableThreatCampaigns.IsNull() || d.DetectionSettings.DisableThreatCampaigns.IsNull() {t.Fatal("explicit marker drift hidden")}
 if d.DetectionSettings.SignatureSelectionSetting.SignatureSettingsByAccuracy.HighAccuracyAction.ValueString()!="SIG_IGNORE" {t.Fatal("explicit signature drift hidden")}
 read(d,map[string]interface{}{"detection_settings":defaults()},false)
 if d.DetectionSettings.EnableSuppression.IsNull() || !d.DetectionSettings.DisableSuppression.IsNull() || d.DetectionSettings.EnableThreatCampaigns.IsNull() || !d.DetectionSettings.DisableThreatCampaigns.IsNull() || d.DetectionSettings.SignatureSelectionSetting.SignatureSettingsByAccuracy.HighAccuracyAction.ValueString()!="SIG_BLOCK" {t.Fatal("policy restoration hidden")}
}
func TestImplicitNondefaultPolicyReadback(t *testing.T) {
 d:=minimal(); wire:=defaults()
 delete(wire,"enable_suppression");wire["disable_suppression"]=map[string]interface{}{}
 delete(wire,"enable_threat_campaigns");wire["disable_threat_campaigns"]=map[string]interface{}{}
 wire["signature_selection_setting"].(map[string]interface{})["signature_settings_by_accuracy"].(map[string]interface{})["low_accuracy_action"]="SIG_BLOCK"
 read(d,map[string]interface{}{"detection_settings":wire},true)
 if d.DetectionSettings.DisableSuppression.IsNull() || d.DetectionSettings.DisableThreatCampaigns.IsNull() || d.DetectionSettings.SignatureSelectionSetting.SignatureSettingsByAccuracy.LowAccuracyAction.ValueString()!="SIG_BLOCK" {t.Fatal("nondefault imported policy lost")}
}
`
	if err := os.WriteFile(filepath.Join(dir, "readback_test.go"), []byte(source), 0600); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", "-count=1", "./internal/"+filepath.Base(dir))
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated WAF readback failed: %v\n%s", err, out)
	}
}
