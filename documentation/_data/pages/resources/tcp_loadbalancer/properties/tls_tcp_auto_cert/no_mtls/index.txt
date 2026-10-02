---
page_title: "tls_tcp_auto_cert.no_mtls"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls tcp auto cert no mtls"], "body_bytes": 1277, "body_sha256": "sha256:4234f3b0be5f4d9f795aa94462a7ecd242e2a9e6a89b35963047b0253a73ca3d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:no_mtls", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert", "path": "documentation/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/no_mtls/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2332303230031210-0302010003210303-0310320200311033-1033332033110010-0002031132312001-3011301111002122-3230213021313202-1110210010312033", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp_auto_cert", "no_mtls"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/no_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert.no_mtls

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [tls_tcp_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/)
- tls_tcp_auto_cert.no_mtls

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
no_mtls = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_tcp_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
