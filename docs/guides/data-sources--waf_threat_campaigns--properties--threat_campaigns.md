---
page_title: "threat_campaigns"
subcategory: ""
description: "threat_campaigns for xcsh_waf_threat_campaigns."
xcsh_docs: {"aliases": [], "body_bytes": 2252, "body_sha256": "sha256:8570d1e3c085206ccf487c26cb10204d133bedf9aef4dacb044a986e97d78c34", "canonical_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "child_ids": [], "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "parent_id": "xcsh-docs:data-sources:waf_threat_campaigns:reference", "path": "docs/guides/data-sources--waf_threat_campaigns--properties--threat_campaigns.md", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["threat_campaigns"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/properties/threat_campaigns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "threat_campaigns for xcsh_waf_threat_campaigns.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# threat_campaigns

Breadcrumbs:

- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md)
- [Property reference](data-sources--waf_threat_campaigns--reference.md)
- threat_campaigns

<a id="section"></a>

Type: `"list"`. Computed.

Threat Campaigns. A list of all supported threat campaigns.

## Direct properties

<a id="schema-threat_campaigns--attack_type"></a>

### attack_type property

Type: `"string"`. Computed.

Attack Type. The Threat Campaign Attack Type.

<a id="schema-threat_campaigns--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. The Threat Campaign Description.

<a id="schema-threat_campaigns--id"></a>

### id property

Type: `"string"`. Computed.

ID. The Threat Campaign ID.

<a id="schema-threat_campaigns--intent"></a>

### intent property

Type: `"string"`. Computed.

Intent. The Threat Campaign Intent.

<a id="schema-threat_campaigns--last_update"></a>

### last_update property

Type: `"string"`. Computed.

Last Update. The Threat Campaign last update time.

<a id="schema-threat_campaigns--malwares"></a>

### malwares property

Type: `["list", "string"]`. Computed.

Malwares. The Threat Campaign Malwares.

<a id="schema-threat_campaigns--name"></a>

### name property

Type: `"string"`. Computed.

Name. The Threat Campaign Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-threat_campaigns--references"></a>

### references property

Type: `["list", "string"]`. Computed.

References. The Threat Campaign References.

<a id="schema-threat_campaigns--risk"></a>

### risk property

Type: `"string"`. Computed.

Risk. The Threat Campaign Risk.

<a id="schema-threat_campaigns--systems"></a>

### systems property

Type: `["list", "string"]`. Computed.

Systems. The Threat Campaign Systems.

## Next pages

- [Property reference](data-sources--waf_threat_campaigns--reference.md)
- [xcsh_waf_threat_campaigns](../data-sources/waf_threat_campaigns.md)
