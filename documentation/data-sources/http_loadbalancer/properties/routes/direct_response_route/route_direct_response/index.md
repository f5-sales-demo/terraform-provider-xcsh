---
page_title: "routes.direct_response_route.route_direct_response"
subcategory: "Load Balancing"
description: "Send this direct response in case of route match action is direct response."
xcsh_docs: {"aliases": ["routes direct response route route direct response"], "body_bytes": 3216, "body_sha256": "sha256:a2fac93aea35acf35733108e71e170fa84a288ff3fe34a27a3332536ddc2612b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:route_direct_response", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route", "path": "documentation/data-sources/http_loadbalancer/properties/routes/direct_response_route/route_direct_response/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2331322212101113-0300120012002100-1021210210302331-3012003202132110-0311320022310330-3332012312231230-3332321122031331-1103300221333021", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "direct_response_route", "route_direct_response"], "schema_version": 1, "sections": [{"aliases": ["routes direct response route route direct response response body encoded"], "anchor": "schema-routes--direct_response_route--route_direct_response--response_body_encoded", "description": "Response body to send. Currently supported URL schemes is string:/// for which message should be encoded in Base64 format. The message can be either plain text or HTML. E.g. \"<p> Access Denied </p>\". Base64 encoded string URL for this is string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:route_direct_response", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "direct_response_route", "route_direct_response", "response_body_encoded"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes direct response route route direct response response code"], "anchor": "schema-routes--direct_response_route--route_direct_response--response_code", "description": "Response code to send.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:route_direct_response", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "direct_response_route", "route_direct_response", "response_code"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/direct_response_route/route_direct_response/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Send this direct response in case of route match action is direct response.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.direct_response_route.route_direct_response

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [routes.direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/direct_response_route/)
- routes.direct_response_route.route_direct_response

<a id="section"></a>

Type: `"single"`. Computed.

Send this direct response in case of route match action is direct response.

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

<a id="schema-routes--direct_response_route--route_direct_response--response_body_encoded"></a>

### response_body_encoded property

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". Base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="schema-routes--direct_response_route--route_direct_response--response_code"></a>

### response_code property

Type: `"number"`. Computed.

Response Code. Response code to send.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```
