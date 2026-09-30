---
page_title: "advertise_on_public"
subcategory: "Load Balancing"
description: "advertise_on_public for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1612, "body_sha256": "sha256:ee9ac7abda1d3ce35d4816dccad0e8db4f49bd0e88d5e0079e19f22d6c1eda43", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:advertise_on_public:public_ip"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_on_public", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/advertise_on_public/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["advertise_on_public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/advertise_on_public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_on_public for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# advertise_on_public

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- advertise_on_public

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_on_public {
  # Configure direct properties listed below.
}
```

## Direct properties

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_on_public/public_ip/): complete subsection reference.

## Next pages

- [advertise_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_on_public/public_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
