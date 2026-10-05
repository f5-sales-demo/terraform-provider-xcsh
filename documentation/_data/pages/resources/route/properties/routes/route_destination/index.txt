---
page_title: "routes.route_destination"
subcategory: ""
description: "List of destination to choose if the route is match."
xcsh_docs: {"aliases": ["routes route destination"], "body_bytes": 13468, "body_sha256": "sha256:f759d962f331bd54cb6a8d9ab129ada3a68e59ade666cee22ce1cce2b1f6b52c", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:route_destination:buffer_policy", "xcsh-docs:resources:route:properties:routes:route_destination:cors_policy", "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy", "xcsh-docs:resources:route:properties:routes:route_destination:destinations", "xcsh-docs:resources:route:properties:routes:route_destination:do_not_retract_cluster", "xcsh-docs:resources:route:properties:routes:route_destination:endpoint_subsets", "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy", "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "xcsh-docs:resources:route:properties:routes:route_destination:retract_cluster", "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "xcsh-docs:resources:route:properties:routes:route_destination:spdy_config", "xcsh-docs:resources:route:properties:routes:route_destination:web_socket_config"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:route_destination", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "documentation/resources/route/properties/routes/route_destination/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310", "registry_path": "docs/guides/resources--route--reference--group-002.md", "relationships": [{"anchor": "schema-routes--route_destination--auto_host_rewrite", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--host_rewrite", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--prefix_rewrite", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:do_not_retract_cluster,retract_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:do_not_retract_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:do_not_retract_cluster,retract_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:retract_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination:RequiredObjectAttributes:destinations", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "route_destination"], "schema_version": 1, "sections": [{"aliases": ["routes route destination auto host rewrite"], "anchor": "schema-routes--route_destination--auto_host_rewrite", "description": "Exclusive with Indicates that during forwarding, the host header will be swapped with the hostname of the upstream host chosen by the cluster.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "auto_host_rewrite"], "syntax": "attribute", "type": "bool"}, {"aliases": ["routes route destination buffer policy"], "anchor": "section", "description": "Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:buffer_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "buffer_policy"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "routes route destination cors policy"], "anchor": "section", "description": "Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre)", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:cors_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "cors_policy"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination csrf policy"], "anchor": "section", "description": "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:csrf_policy:disabled", "type": "conflicts"}], "schema_path": ["routes", "route_destination", "csrf_policy"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination destinations"], "anchor": "section", "description": "When requests have to distributed among multiple upstream clusters, multiple destinations are configured, each having its own cluster and weight. Traffic is distributed among clusters based on the weight configured. Example: destinations: - cluster: - kind: F5 xc.vega.cfg.adc.cluster.object uid: cluster-1 weight: 20 -", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.destinations:RequiredListObjectAttributes:cluster", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations:cluster", "type": "requires"}], "schema_path": ["routes", "route_destination", "destinations"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination do not retract cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:do_not_retract_cluster", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "do_not_retract_cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination endpoint subsets"], "anchor": "section", "description": "Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer Labels field of endpoint object's metadata is used for subset matching. For endpoint's which", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:endpoint_subsets", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "endpoint_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination hash policy"], "anchor": "section", "description": "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--route_destination--hash_policy--header_name", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:cookie,header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--hash_policy--header_name", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:header_name,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--hash_policy--source_ip", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:cookie,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--hash_policy--source_ip", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:header_name,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:cookie,header_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.hash_policy:ConflictingListObjectAttributes:cookie,source_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:hash_policy:cookie", "type": "conflicts"}], "schema_path": ["routes", "route_destination", "hash_policy"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination host rewrite"], "anchor": "schema-routes--route_destination--host_rewrite", "description": "Exclusive with Indicates that during forwarding, the host header will be swapped with this value.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "host_rewrite"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes route destination mirror policy"], "anchor": "section", "description": "MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is \"fire and forget\", meaning it will not wait for the shadow cluster to respond before returning the response from the primary cluster. All normal statistics are collected for the shadow cluster making this feature useful for", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.mirror_policy:RequiredObjectAttributes:cluster", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:mirror_policy:cluster", "type": "requires"}], "schema_path": ["routes", "route_destination", "mirror_policy"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination prefix rewrite"], "anchor": "schema-routes--route_destination--prefix_rewrite", "description": "Exclusive with prefix_rewrite indicates that during forwarding, the matched prefix (or path) should be swapped with its value. When using regex path matching, the entire path (not including the query string) will be swapped with this value. This option allows application URLs to be rooted at a different path from", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "prefix_rewrite"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes route destination priority"], "anchor": "schema-routes--route_destination--priority", "description": "Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["DEFAULT", "HIGH"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "priority"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes route destination query params"], "anchor": "section", "description": "Handling of incoming query parameters in simple route.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--route_destination--query_params--replace_params", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--query_params--replace_params", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:retain_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination.query_params:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:query_params:retain_all_params", "type": "conflicts"}], "schema_path": ["routes", "route_destination", "query_params"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination regex rewrite"], "anchor": "section", "description": "RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "regex_rewrite"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination retract cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retract_cluster", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "retract_cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes route destination retry policy"], "anchor": "section", "description": "Retry policy configuration for route destination.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--route_destination--retry_policy--retry_condition", "enforcement": "provider-schema", "group": "routes.route_destination.retry_policy:RequiredObjectAttributes:retry_condition", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:retry_policy", "type": "requires"}], "schema_path": ["routes", "route_destination", "retry_policy"], "syntax": "block", "type": "object"}, {"aliases": ["routes route destination spdy config"], "anchor": "section", "description": "Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1' Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'SPDY/3.1'", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:spdy_config", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "spdy_config"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "routes route destination timeout"], "anchor": "schema-routes--route_destination--timeout", "description": "Specifies the timeout for the route in milliseconds. This timeout includes all retries. For server side streaming, configure this field with higher value or leave it un-configured for infinite timeout.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "route_destination", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["routes route destination web socket config"], "anchor": "section", "description": "Configuration to allow Websocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination:web_socket_config", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "route_destination", "web_socket_config"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/route_destination/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "List of destination to choose if the route is match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_destination

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- routes.route_destination

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of destination to choose if the route is match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("destinations"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("do_not_retract_cluster",
    "retract_cluster"),
  validators.ConflictingObjectAttributes("prefix_rewrite",
    "regex_rewrite")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"host_rewrite\"]",
  "x-ves-oneof-field-route_destination_rewrite": "[\"prefix_rewrite\",\"regex_rewrite\"]"
}
```

Terraform syntax:

```terraform
route_destination {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--route_destination--auto_host_rewrite"></a>

### auto_host_rewrite property

Type: `"bool"`. Optional.

Exclusive with \[host\_rewrite\] Indicates that during forwarding, the host header will be swapped
with the hostname of the upstream host chosen by the cluster.

Upstream description:

Exclusive with \[host\_rewrite\] Indicates that during forwarding, the host header will be swapped
with the hostname of the upstream host chosen by the cluster.

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

- [buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/buffer_policy/): complete subsection reference.

- [cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/): complete subsection reference.

- [csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/): complete subsection reference.

- [destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/): complete subsection reference.

- [do_not_retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/do_not_retract_cluster/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/endpoint_subsets/): complete subsection reference.

- [hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/): complete subsection reference.

<a id="schema-routes--route_destination--host_rewrite"></a>

### host_rewrite property

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite\] Indicates that during forwarding, the host header will be
swapped with this value.

Upstream description:

Exclusive with \[auto\_host\_rewrite\] Indicates that during forwarding, the host header will be
swapped with this value.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/): complete subsection reference.

<a id="schema-routes--route_destination--prefix_rewrite"></a>

### prefix_rewrite property

Type: `"string"`. Optional.

Exclusive with \[regex\_rewrite\] prefix\_rewrite indicates that during forwarding, the matched
prefix (or path) should be swapped with its value. When using regex path matching, the entire path
(not including the query string) will be swapped with this value. This option allows application
URLs to..

Upstream description:

Exclusive with \[regex\_rewrite\] prefix\_rewrite indicates that during forwarding, the matched
prefix (or path) should be swapped with its value. When using regex path matching, the entire path
(not including the query string) will be swapped with this value. This option allows application
URLs to be rooted at a different path from those exposed at the reverse proxy layer.

Example : gcSpec: routes: &#8203;- match: &#8203;- headers: \[\] path: prefix : /register/
query\_params: \[\] &#8203;- headers: \[\] path: prefix: /register query\_params: \[\]
routeDestination: prefixRewrite: "/" destinations: &#8203;- cluster: &#8203;- kind: cluster.object
uid: cluster-1

Having above entries in the config, requests to /register will be stripped to /, while requests to
/register/public will be stripped to /public.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-routes--route_destination--priority"></a>

### priority property

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DEFAULT","HIGH"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/): complete subsection reference.

- [regex_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/regex_rewrite/): complete subsection reference.

- [retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retract_cluster/): complete subsection reference.

- [retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/): complete subsection reference.

- [spdy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/spdy_config/): complete subsection reference.

<a id="schema-routes--route_destination--timeout"></a>

### timeout property

Type: `"number"`. Optional.

Specifies the timeout for the route in milliseconds. This timeout includes all retries. For server
side streaming, configure this field with higher value or leave it un-configured for infinite
timeout.

Upstream description:

Specifies the timeout for the route in milliseconds. This timeout includes all retries. For server
side streaming, configure this field with higher value or leave it un-configured for infinite
timeout.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1800000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

- [web_socket_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/web_socket_config/): complete subsection reference.

## Next pages

- [routes.route_destination.buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/buffer_policy/)
- [routes.route_destination.cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/)
- [routes.route_destination.csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/)
- [routes.route_destination.destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/)
- [routes.route_destination.do_not_retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/do_not_retract_cluster/)
- [routes.route_destination.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/endpoint_subsets/)
- [routes.route_destination.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/)
- [routes.route_destination.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/)
- [routes.route_destination.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/)
- [routes.route_destination.regex_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/regex_rewrite/)
- [routes.route_destination.retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retract_cluster/)
- [routes.route_destination.retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/)
- [routes.route_destination.spdy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/spdy_config/)
- [routes.route_destination.web_socket_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/web_socket_config/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
