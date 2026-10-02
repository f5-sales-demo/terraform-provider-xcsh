---
page_title: "slow_ddos_mitigation.disable_request_timeout"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["duration", "operation timeout", "slow ddos mitigation disable request timeout"], "body_bytes": 1349, "body_sha256": "sha256:1aa704c9318c948e080003384f8d91fa364d9cccaf1e57a3460d0819d23e1345", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation:disable_request_timeout", "parent_id": "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation", "path": "documentation/resources/virtual_host/properties/slow_ddos_mitigation/disable_request_timeout/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0302112002310101-2021323133201110-2302121112300222-2022200302212321-1020202132131320-1021002113020201-2320232221231103-3212312020311302", "registry_path": "docs/guides/resources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["slow_ddos_mitigation", "disable_request_timeout"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/slow_ddos_mitigation/disable_request_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slow_ddos_mitigation.disable_request_timeout

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/slow_ddos_mitigation/)
- slow_ddos_mitigation.disable_request_timeout

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

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
disable_request_timeout = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/slow_ddos_mitigation/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
