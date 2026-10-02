---
page_title: "access_info.tls_config.disable_session_key_caching"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["access info tls config disable session key caching"], "body_bytes": 1593, "body_sha256": "sha256:4c82412c8cef288205c25722dad996f4a8b49fc5863f377d957940b9bb01db11", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_session_key_caching", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "path": "documentation/resources/secret_management_access/properties/access_info/tls_config/disable_session_key_caching/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2203002320210003-3112320211113020-1231121133010311-1233120131013313-3001311220012032-3222010112020103-2122300120312001-3113010110113312", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "disable_session_key_caching"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/disable_session_key_caching/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.disable_session_key_caching

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/)
- access_info.tls_config.disable_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
