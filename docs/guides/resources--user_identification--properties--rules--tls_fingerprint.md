---
page_title: "rules.tls_fingerprint"
subcategory: ""
description: "rules.tls_fingerprint for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1019, "body_sha256": "sha256:24a75bb920bf8a24e760c12141698c35a33421a2be25d1e9ebd32ba314f6c84c", "canonical_id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "child_ids": [], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules:tls_fingerprint", "parent_id": "xcsh-docs:resources:user_identification:properties:rules", "path": "docs/guides/resources--user_identification--properties--rules--tls_fingerprint.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "tls_fingerprint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/tls_fingerprint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.tls_fingerprint for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.tls_fingerprint

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md)
- [Property reference](resources--user_identification--reference.md)
- [rules](resources--user_identification--properties--rules.md)
- rules.tls_fingerprint

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for tls fingerprint.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
tls_fingerprint = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules](resources--user_identification--properties--rules.md)
- [xcsh_user_identification](../resources/user_identification.md)
