---
page_title: "attack_signatures"
subcategory: ""
description: "attack_signatures for xcsh_waf_attack_signatures."
xcsh_docs: {"aliases": [], "body_bytes": 2098, "body_sha256": "sha256:74994177321db9bec7108793dcd8ffd3cdd4e8248df11991f338f917fb5b40cc", "canonical_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "child_ids": [], "collection_id": "xcsh-docs:data-sources:waf_attack_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "parent_id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "path": "docs/guides/data-sources--waf_attack_signatures--properties--attack_signatures.md", "provider_name": "waf_attack_signatures", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["attack_signatures"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_attack_signatures/properties/attack_signatures/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "attack_signatures for xcsh_waf_attack_signatures.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# attack_signatures

Breadcrumbs:

- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md)
- [Property reference](data-sources--waf_attack_signatures--reference.md)
- attack_signatures

<a id="section"></a>

Type: `"list"`. Computed.

List of all supported attack signatures.

## Direct properties

<a id="schema-attack_signatures--accuracy"></a>

### accuracy property

Type: `"string"`. Computed.

Accuracy. The Signature Accuracy.

<a id="schema-attack_signatures--applies_to"></a>

### applies_to property

Type: `"string"`. Computed.

Applies To. The Signature Applies to.

<a id="schema-attack_signatures--attack_type"></a>

### attack_type property

Type: `"string"`. Computed.

Attack Type. The Signature Attack Type.

<a id="schema-attack_signatures--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. The Signature Description.

<a id="schema-attack_signatures--id"></a>

### id property

Type: `"string"`. Computed.

ID. The Signature ID.

<a id="schema-attack_signatures--last_update"></a>

### last_update property

Type: `"string"`. Computed.

Last Update. The Signature last update time.

<a id="schema-attack_signatures--name"></a>

### name property

Type: `"string"`. Computed.

Name. The Signature Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-attack_signatures--references"></a>

### references property

Type: `["list", "string"]`. Computed.

References. The Signature References.

<a id="schema-attack_signatures--risk"></a>

### risk property

Type: `"string"`. Computed.

Risk. The Signature Risk.

<a id="schema-attack_signatures--systems"></a>

### systems property

Type: `["list", "string"]`. Computed.

Systems. The Signature Systems.

## Next pages

- [Property reference](data-sources--waf_attack_signatures--reference.md)
- [xcsh_waf_attack_signatures](../data-sources/waf_attack_signatures.md)
