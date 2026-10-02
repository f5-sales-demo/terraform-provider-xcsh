---
page_title: "ingress_gw.allowed_vip_port.custom_ports"
subcategory: "Infrastructure"
description: "List of Custom port."
xcsh_docs: {"aliases": ["ingress gw allowed vip port custom ports"], "body_bytes": 2873, "body_sha256": "sha256:dba85bb31d13876e652d1e79f264bd4f2169ef759f02c93104f035355b9fe116", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port", "path": "documentation/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/custom_ports/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3030130100031031-3110330103102133-2211221032232332-3121213010313222-1122011110303122-3201121130233003-1310312200330102-0332130013110010", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-004.md", "relationships": [{"anchor": "schema-ingress_gw--allowed_vip_port--custom_ports--port_ranges", "enforcement": "provider-schema", "group": "ingress_gw.allowed_vip_port.custom_ports:RequiredObjectAttributes:port_ranges", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw", "allowed_vip_port", "custom_ports"], "schema_version": 1, "sections": [{"aliases": ["port ranges"], "anchor": "schema-ingress_gw--allowed_vip_port--custom_ports--port_ranges", "description": "Port Ranges.", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_gw:allowed_vip_port:custom_ports", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_gw", "allowed_vip_port", "custom_ports", "port_ranges"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/custom_ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of Custom port.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.allowed_vip_port.custom_ports

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/)
- [ingress_gw.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/)
- ingress_gw.allowed_vip_port.custom_ports

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
```

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
custom_ports {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ingress_gw--allowed_vip_port--custom_ports--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

## Next pages

- [ingress_gw.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/ingress_gw/allowed_vip_port/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
