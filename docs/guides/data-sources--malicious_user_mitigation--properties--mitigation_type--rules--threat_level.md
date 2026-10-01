---
page_title: "mitigation_type.rules.threat_level"
subcategory: ""
description: "mitigation_type.rules.threat_level for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 2011, "body_sha256": "sha256:81903112647c2415e32b07ade9b0f17bc8d49bdf6a339ea45922f0eb86d16783", "canonical_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "child_ids": ["xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:high", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:low", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level:medium"], "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:threat_level", "parent_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules", "path": "docs/guides/data-sources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mitigation_type", "rules", "threat_level"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/threat_level/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mitigation_type.rules.threat_level for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules.threat_level

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md)
- [Property reference](data-sources--malicious_user_mitigation--reference.md)
- [mitigation_type](data-sources--malicious_user_mitigation--properties--mitigation_type.md)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--properties--mitigation_type--rules.md)
- mitigation_type.rules.threat_level

<a id="section"></a>

Type: `"single"`. Computed.

Threat level estimated for each user based on the user's activity and reputation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-threat_level": "[\"high\",\"low\",\"medium\"]"
}
```

## Direct properties

- [high](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--high.md): complete subsection reference.

- [low](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--low.md): complete subsection reference.

- [medium](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--medium.md): complete subsection reference.

## Next pages

- [mitigation_type.rules.threat_level.high](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--high.md)
- [mitigation_type.rules.threat_level.low](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--low.md)
- [mitigation_type.rules.threat_level.medium](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--threat_level--medium.md)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--properties--mitigation_type--rules.md)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md)
