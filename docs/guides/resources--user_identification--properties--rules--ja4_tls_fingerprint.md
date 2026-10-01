---
page_title: "rules.ja4_tls_fingerprint"
subcategory: ""
description: "rules.ja4_tls_fingerprint for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1035, "body_sha256": "sha256:0b1072d2854584f09f17404d810e6c517c7af96f6e5d6069ef558c3c3638f1a7", "canonical_id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "child_ids": [], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules:ja4_tls_fingerprint", "parent_id": "xcsh-docs:resources:user_identification:properties:rules", "path": "docs/guides/resources--user_identification--properties--rules--ja4_tls_fingerprint.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "ja4_tls_fingerprint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/ja4_tls_fingerprint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.ja4_tls_fingerprint for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ja4_tls_fingerprint

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md)
- [Property reference](resources--user_identification--reference.md)
- [rules](resources--user_identification--properties--rules.md)
- rules.ja4_tls_fingerprint

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ja4 tls fingerprint.

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
ja4_tls_fingerprint = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules](resources--user_identification--properties--rules.md)
- [xcsh_user_identification](../resources/user_identification.md)
