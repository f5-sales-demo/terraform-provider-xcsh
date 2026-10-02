---
page_title: "rules.ip_and_tls_fingerprint"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules ip and tls fingerprint"], "body_bytes": 1271, "body_sha256": "sha256:d1de0a015f9b62cba66efcff240da5ec60ac68b639ba17f80d9263f033e656de", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "parent_id": "xcsh-docs:resources:user_identification:properties:rules", "path": "documentation/resources/user_identification/properties/rules/ip_and_tls_fingerprint/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2311120321201223-3201102323200101-2023020101303322-1232011222231021-1121223301223211-1231030121302331-3123211030232003-0323001332231000", "registry_path": "docs/guides/resources--user_identification--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "ip_and_tls_fingerprint"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/ip_and_tls_fingerprint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ip_and_tls_fingerprint

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/)
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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/)
- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
