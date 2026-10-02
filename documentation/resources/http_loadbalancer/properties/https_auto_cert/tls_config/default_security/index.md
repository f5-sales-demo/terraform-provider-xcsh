---
page_title: "https_auto_cert.tls_config.default_security"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["https auto cert tls config default security"], "body_bytes": 1498, "body_sha256": "sha256:ab409e56c7e6d2a858062bc60068044760eb3862dd748a164b392b5fa8d61432", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:tls_config:default_security", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:tls_config", "path": "documentation/resources/http_loadbalancer/properties/https_auto_cert/tls_config/default_security/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3223202003303223-2033112010233130-1300011333021230-0033120002012013-3203210120313021-2202032331002232-3011212300023201-3222120202213303", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_auto_cert", "tls_config", "default_security"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https_auto_cert/tls_config/default_security/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert.tls_config.default_security

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https_auto_cert/)
- [https_auto_cert.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https_auto_cert/tls_config/)
- https_auto_cert.tls_config.default_security

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
default_security = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https_auto_cert.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https_auto_cert/tls_config/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
