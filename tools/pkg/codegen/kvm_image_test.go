package codegen

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const syntheticKVMImageContract = `{
 "availability":"evidence_backed","enforcement":"required","strategy":"site_uid_os_image","namespace":"system",
 "terraform_data_source":"site_image","replaces_operation":"ves.io.schema.registration.CustomAPI.GetImageDownloadUrl",
 "configuration_get":{"method":"GET","path":"/api/config/namespaces/system/securemesh_site_v2s/{site_name}","operation_id":"ves.io.schema.views.securemesh_site_v2.API.Get"},
 "configuration_list":{"method":"GET","path":"/api/config/namespaces/system/securemesh_site_v2s","operation_id":"ves.io.schema.views.securemesh_site_v2.API.List"},
 "site_list":{"method":"GET","path":"/api/config/namespaces/system/sites","operation_id":"ves.io.schema.site.API.List"},
 "identity":{"configuration_name":"items[].name","configuration_uid":"items[].uid","site_uid":"items[].uid","owner_kind":"items[].owner_view.kind","required_owner_kind":"securemesh_site_v2","owner_uid":"items[].owner_view.uid","request_uid":"site_uid","configuration_matches":1,"owner_matches":1},
 "query":{"method":"POST","path":"/api/maurice/software_os_version","operation_id":"ves.io.schema.virtual_appliance.SoftwareVersionOsImageCustomApi.GetImage","request_schema":"virtual_applianceGetImageRequest","response_schema":"virtual_applianceGetImageResponse","request_field":"uids","request_cardinality":1,"side_effects":"none"},
 "response":{"mapping":"images","mapping_key":"site_uid","download_url":"download_image_link","image_name":"copy_image_name","md5":"image_md5_sum","error":"error_description"},
 "validation":{"caller_supplied_uid":false,"static_fallback":false,"require_current_owner_mapping":true,"required_platform_field":"spec.kvm","reject_nonempty_error":true,"require_complete_image_fields":true,"download_scheme":"https","download_hosts":["downloads.volterra.io"],"md5_pattern":"^[0-9a-fA-F]{32}$","verify_artifact_checksum":true,"boot_acceptance_required":true},
 "provenance":{"source_issue":"f5-sales-demo/api-specs-enriched#1805","source_commit":"5c33dcf51eb8bee24e2ca1856e32cb5fef2a286f","upstream_spec_sha256":"b36cd80bf6394741bdc8ba88ef175251129ae12be8b8f2d6eed2676227f704ef","receipt_path":"config/evidence/kvm-site-image-resolution-20260919.json","receipt_sha256":"ca100108cda383ca0a5cd99bec804933e179b5c4a23b45d907d2ea1de520edf7"}
}`

func TestKVMImageContractRejectsIncompleteOrInventedSemantics(t *testing.T) {
	var valid map[string]any
	if err := json.Unmarshal([]byte(syntheticKVMImageContract), &valid); err != nil {
		t.Fatal(err)
	}
	if err := validateKVMImageResolution(valid); err != nil {
		t.Fatal(err)
	}
	for key := range valid {
		t.Run("missing "+key, func(t *testing.T) {
			var broken map[string]any
			_ = json.Unmarshal([]byte(syntheticKVMImageContract), &broken)
			delete(broken, key)
			if validateKVMImageResolution(broken) == nil {
				t.Fatal("incomplete contract accepted")
			}
		})
	}
	for _, pair := range [][2]string{{`"request_uid":"site_uid"`, `"request_uid":"configuration_uid"`}, {`"static_fallback":false`, `"static_fallback":true`}, {`"caller_supplied_uid":false`, `"caller_supplied_uid":true`}, {`"required_platform_field":"spec.kvm"`, `"required_platform_field":"spec.aws"`}, {`"receipt_sha256":"ca100108cda383ca0a5cd99bec804933e179b5c4a23b45d907d2ea1de520edf7"`, `"receipt_sha256":"unknown"`}} {
		var broken map[string]any
		_ = json.Unmarshal([]byte(strings.Replace(syntheticKVMImageContract, pair[0], pair[1], 1)), &broken)
		if validateKVMImageResolution(broken) == nil {
			t.Error("invented semantics accepted")
		}
	}
}

func TestKVMImageSuppressedOperationsExcludeRawCallerUIDSurface(t *testing.T) {
	contract := []byte(`{"providers":{"kvm":{"image_resolution":` + syntheticKVMImageContract + `}}}`)
	got, err := SMSv2ImageSuppressedOperations(contract)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]struct{}{
		"ves.io.schema.registration.CustomAPI.GetImageDownloadUrl":                 {},
		"ves.io.schema.virtual_appliance.SoftwareVersionOsImageCustomApi.GetImage": {},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("suppressed operations = %#v, want %#v", got, want)
	}
}

func TestGenerateKVMImageDataSourceIsCleanBreakAndDeterministic(t *testing.T) {
	var image map[string]any
	_ = json.Unmarshal([]byte(syntheticKVMImageContract), &image)
	dir := t.TempDir()
	if err := generateKVMImageDataSource(image, dir); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "site_image_data_source.go")
	first, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := generateKVMImageDataSource(image, dir); err != nil {
		t.Fatal(err)
	}
	second, _ := os.ReadFile(path)
	if string(first) != string(second) {
		t.Fatal("nondeterministic image generation")
	}
	for _, want := range []string{"Code generated", "Source: F5 XC enriched API response-operation contract", "ResolveKVMImage", `tfsdk:"site_name"`, `tfsdk:"image_md5_sum"`, "Sensitive: true", "/api/maurice/software_os_version"} {
		if !strings.Contains(string(first), want) {
			t.Errorf("generated image resolver missing %q", want)
		}
	}
	for _, absent := range []string{"get-image-download-url", "ProviderRef", "image_md5_download_url", `tfsdk:"uid"`} {
		if strings.Contains(string(first), absent) {
			t.Errorf("generated image resolver retained obsolete surface %q", absent)
		}
	}
}
