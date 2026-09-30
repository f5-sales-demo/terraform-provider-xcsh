---
page_title: "ingress_egress_gw.no_network_policy"
subcategory: "Infrastructure"
description: "ingress_egress_gw.no_network_policy for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 943, "body_sha256": "sha256:eebc50c1c06cf7e4e33de051e94fde4134658c61bfa678414b6c3b01d763d85d", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_network_policy", "child_ids": [], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_network_policy", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "path": "docs/guides/resources--gcp_vpc_site--properties--ingress_egress_gw--no_network_policy.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "no_network_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/no_network_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.no_network_policy for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.no_network_policy

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [ingress_egress_gw](resources--gcp_vpc_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.no_network_policy

<a id="section"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ingress_egress_gw](resources--gcp_vpc_site--properties--ingress_egress_gw.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
