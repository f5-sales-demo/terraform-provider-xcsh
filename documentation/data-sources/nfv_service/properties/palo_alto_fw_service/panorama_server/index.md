---
page_title: "palo_alto_fw_service.panorama_server"
subcategory: ""
description: "Panorama Server Type."
xcsh_docs: {"aliases": ["palo alto fw service panorama server"], "body_bytes": 4173, "body_sha256": "sha256:2328bf87c3fe5b8996bff65fe682c544bf74e006c96c4903e988c733ae0c3276", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3201221223132300-0002223112302331-2310200230113210-1122311213321321-0230122201303030-0321223210111331-1133312200221010-1312222322121200", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "panorama_server"], "schema_version": 1, "sections": [{"aliases": ["authorization key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "panorama_server", "authorization_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["device group name"], "anchor": "schema-palo_alto_fw_service--panorama_server--device_group_name", "description": "Device Group Name.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "panorama_server", "device_group_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["server"], "anchor": "schema-palo_alto_fw_service--panorama_server--server", "description": "Panorama Server Address to which the firewall should connect to.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "panorama_server", "server"], "syntax": "attribute", "type": "string"}, {"aliases": ["template stack name"], "anchor": "schema-palo_alto_fw_service--panorama_server--template_stack_name", "description": "Template Stack Name.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "panorama_server", "template_stack_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Panorama Server Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.panorama_server

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- palo_alto_fw_service.panorama_server

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for panorama server.

Upstream description:

Panorama Server Type.

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

## Direct properties

- [authorization_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/): complete subsection reference.

<a id="schema-palo_alto_fw_service--panorama_server--device_group_name"></a>

### device_group_name property

Type: `"string"`. Computed.

Device Group Name. Device Group Name.

Upstream description:

Device Group Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-palo_alto_fw_service--panorama_server--server"></a>

### server property

Type: `"string"`. Computed.

Panorama Server Address to which the firewall should connect to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-palo_alto_fw_service--panorama_server--template_stack_name"></a>

### template_stack_name property

Type: `"string"`. Computed.

Template stack name. Template Stack Name.

Upstream description:

Template Stack Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

## Next pages

- [palo_alto_fw_service.panorama_server.authorization_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
