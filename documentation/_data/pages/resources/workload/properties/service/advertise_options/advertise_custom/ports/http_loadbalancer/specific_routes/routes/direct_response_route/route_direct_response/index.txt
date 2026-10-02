---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response"
subcategory: "Container"
description: "Send this direct response in case of route match action is direct response."
xcsh_docs: {"aliases": ["service advertise options advertise custom ports http loadbalancer specific routes routes direct response route route direct response"], "body_bytes": 6164, "body_sha256": "sha256:22c64c06a171543b6d5385dca065bcbc4a29c6583770306cd0e2e4527ad13fb1", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:route_direct_response", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/direct_response_route/route_direct_response/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0013230300013321-2022130123233200-1323122002332002-1233012001231310-3023020323223100-3202010020323203-3302111201322232-0013031221012013", "registry_path": "docs/guides/resources--workload--reference--group-008.md", "relationships": [{"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--route_direct_response--response_code", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response:RequiredObjectAttributes:response_code", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:route_direct_response", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "route_direct_response"], "schema_version": 1, "sections": [{"aliases": ["response body encoded"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--route_direct_response--response_body_encoded", "description": "Response body to send. Currently supported URL schemes is string:/// for which message should be encoded in Base64 format. The message can be either plain text or HTML. E.g. \"<p> Access Denied </p>\". Base64 encoded string URL for this is string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:route_direct_response", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "route_direct_response", "response_body_encoded"], "syntax": "attribute", "type": "string"}, {"aliases": ["response code"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--route_direct_response--response_code", "description": "Response code to send.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:route_direct_response", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "route_direct_response", "response_code"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/direct_response_route/route_direct_response/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Send this direct response in case of route match action is direct response.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/direct_response_route/)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Send this direct response in case of route match action is direct response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
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
route_direct_response {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--route_direct_response--response_body_encoded"></a>

### response_body_encoded property

Type: `"string"`. Optional.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML.

Upstream description:

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". Base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--route_direct_response--response_code"></a>

### response_code property

Type: `"number"`. Optional.

Response Code. Response code to send.

Upstream description:

Response code to send.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(100, 599),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/direct_response_route/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
