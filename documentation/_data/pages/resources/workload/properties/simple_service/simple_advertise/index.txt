---
page_title: "simple_service.simple_advertise"
subcategory: "Container"
description: "Advertise OPTIONS for Simple Service."
xcsh_docs: {"aliases": ["simple service simple advertise"], "body_bytes": 5110, "body_sha256": "sha256:dafac6fd4ef3646b369df2a94a1d4fd19cf92b9a1d438999e76b327fc461cf48", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "parent_id": "xcsh-docs:resources:workload:properties:simple_service", "path": "documentation/resources/workload/properties/simple_service/simple_advertise/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3310302312301012-0213323302330302-2012000032301123-3011230301130001-2021102001021321-2321111011331112-1000322300330133-3212013211230202", "registry_path": "docs/guides/resources--workload--reference--group-018.md", "relationships": [{"anchor": "schema-simple_service--simple_advertise--domains", "enforcement": "provider-schema", "group": "simple_service.simple_advertise:RequiredObjectAttributes:domains,service_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "type": "requires"}, {"anchor": "schema-simple_service--simple_advertise--service_port", "enforcement": "provider-schema", "group": "simple_service.simple_advertise:RequiredObjectAttributes:domains,service_port", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "simple_advertise"], "schema_version": 1, "sections": [{"aliases": ["simple service simple advertise domains"], "anchor": "schema-simple_service--simple_advertise--domains", "description": "A list of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are supported in the suffix or prefix form Supported Domains and search order: 1. Exact Domain names: www.example.com. 2. Domains starting with a Wildcard: *.example.com. Not supported Domains: - Just a Wildcard: * - A", "document_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "simple_advertise", "domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["simple service simple advertise service port"], "anchor": "schema-simple_service--simple_advertise--service_port", "description": "Service port to advertise on Internet via HTTP loadbalancer using port 80.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:simple_advertise", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "simple_advertise", "service_port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/simple_advertise/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Advertise OPTIONS for Simple Service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.simple_advertise

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/)
- simple_service.simple_advertise

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple advertise.

Additional upstream details:

Advertise OPTIONS for Simple Service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains",
    "service_port")}
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
simple_advertise {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-simple_service--simple_advertise--domains"></a>

### domains property

Type: `["list", "string"]`. Optional.

A list of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Load Balancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-simple_service--simple_advertise--service_port"></a>

### service_port property

Type: `"number"`. Optional.

Service port to advertise on Internet via HTTP loadbalancer using port 80.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1024, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```
