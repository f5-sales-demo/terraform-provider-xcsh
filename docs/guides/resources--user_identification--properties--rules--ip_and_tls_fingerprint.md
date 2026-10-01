---
page_title: "rules.ip_and_tls_fingerprint"
subcategory: ""
description: "rules.ip_and_tls_fingerprint for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1014, "body_sha256": "sha256:ac87754cce51cf3e77997e494ae1ea4e219f8e370c1b160504ed9e433acf5cd5", "canonical_id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "child_ids": [], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "parent_id": "xcsh-docs:resources:user_identification:properties:rules", "path": "docs/guides/resources--user_identification--properties--rules--ip_and_tls_fingerprint.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "ip_and_tls_fingerprint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/ip_and_tls_fingerprint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.ip_and_tls_fingerprint for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ip_and_tls_fingerprint

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md)
- [Property reference](resources--user_identification--reference.md)
- [rules](resources--user_identification--properties--rules.md)
- rules.ip_and_tls_fingerprint

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
ip_and_tls_fingerprint = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules](resources--user_identification--properties--rules.md)
- [xcsh_user_identification](../resources/user_identification.md)
