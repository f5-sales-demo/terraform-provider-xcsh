---
page_title: "https.tls_parameters.use_mtls.xfcc_options"
subcategory: "Load Balancing"
description: "X-Forwarded-Client-Cert header elements to be added to requests."
xcsh_docs: {"aliases": ["https tls parameters use mtls xfcc options"], "body_bytes": 2620, "body_sha256": "sha256:870c1d12f4c1c1ea11f3d708082b647552f8e7520b125171541d7cc81d4a0305", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls:xfcc_options", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls", "path": "documentation/data-sources/http_loadbalancer/properties/https/tls_parameters/use_mtls/xfcc_options/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1322023323220103-0120111112033232-3301100122101000-1022023102030303-3333300212330203-2221222021232123-2032120200322113-2021012121033102", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-019.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "tls_parameters", "use_mtls", "xfcc_options"], "schema_version": 1, "sections": [{"aliases": ["xfcc header elements"], "anchor": "schema-https--tls_parameters--use_mtls--xfcc_options--xfcc_header_elements", "description": "X-Forwarded-Client-Cert header elements to be added to requests.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:https:tls_parameters:use_mtls:xfcc_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "tls_parameters", "use_mtls", "xfcc_options", "xfcc_header_elements"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/https/tls_parameters/use_mtls/xfcc_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "X-Forwarded-Client-Cert header elements to be added to requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters.use_mtls.xfcc_options

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/)
- [https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/)
- [https.tls_parameters.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/use_mtls/)
- https.tls_parameters.use_mtls.xfcc_options

<a id="section"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="schema-https--tls_parameters--use_mtls--xfcc_options--xfcc_header_elements"></a>

### xfcc_header_elements property

Type: `["list", "string"]`. Computed.

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

- [https.tls_parameters.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/https/tls_parameters/use_mtls/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
