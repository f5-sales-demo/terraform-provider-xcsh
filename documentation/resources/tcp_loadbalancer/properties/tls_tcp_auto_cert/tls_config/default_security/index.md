---
page_title: "tls_tcp_auto_cert.tls_config.default_security"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls tcp auto cert tls config default security"], "body_bytes": 1506, "body_sha256": "sha256:e19bd6c567bca4122cad623c61afd2d6d0a7b502c2ed36d4f349b9ce28996f6c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config:default_security", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert:tls_config", "path": "documentation/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/default_security/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1200213310211002-1313301303222331-2312132112130122-1300233203000230-3222000322331122-1023012300132233-2011131302132312-2322000302023220", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp_auto_cert", "tls_config", "default_security"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/default_security/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp_auto_cert.tls_config.default_security

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [tls_tcp_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/)
- [tls_tcp_auto_cert.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/)
- tls_tcp_auto_cert.tls_config.default_security

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

- [tls_tcp_auto_cert.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp_auto_cert/tls_config/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
