---
page_title: "https_management.advertise_on_slo_vip.use_mtls.xfcc_options"
subcategory: ""
description: "X-Forwarded-Client-Cert header elements to be added to requests."
xcsh_docs: {"aliases": ["https management advertise on slo vip use mtls xfcc options"], "body_bytes": 3000, "body_sha256": "sha256:b0f14cba8b209ff03f25f8a28c0d879216b39c4bae289a36973b1c29b1a97c5b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls:xfcc_options", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls", "path": "documentation/resources/nfv_service/properties/https_management/advertise_on_slo_vip/use_mtls/xfcc_options/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3223112210013103-1311113030331010-3121321332011303-2102323201001131-3220022230313220-2312300022222123-0100011211310032-2231110110020003", "registry_path": "docs/guides/resources--nfv_service--reference--group-003.md", "relationships": [{"anchor": "schema-https_management--advertise_on_slo_vip--use_mtls--xfcc_options--xfcc_header_elements", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_vip.use_mtls.xfcc_options:RequiredObjectAttributes:xfcc_header_elements", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls:xfcc_options", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_slo_vip", "use_mtls", "xfcc_options"], "schema_version": 1, "sections": [{"aliases": ["https management advertise on slo vip use mtls xfcc options xfcc header elements"], "anchor": "schema-https_management--advertise_on_slo_vip--use_mtls--xfcc_options--xfcc_header_elements", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:use_mtls:xfcc_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_vip", "use_mtls", "xfcc_options", "xfcc_header_elements"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_slo_vip/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "X-Forwarded-Client-Cert header elements to be added to requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_vip.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/)
- [https_management.advertise_on_slo_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_vip/)
- [https_management.advertise_on_slo_vip.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_vip/use_mtls/)
- https_management.advertise_on_slo_vip.use_mtls.xfcc_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-https_management--advertise_on_slo_vip--use_mtls--xfcc_options--xfcc_header_elements"></a>

### xfcc_header_elements property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

## Next pages

- [https_management.advertise_on_slo_vip.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_vip/use_mtls/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
