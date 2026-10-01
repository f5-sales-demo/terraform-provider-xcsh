---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 27102, "body_sha256": "sha256:07af7e5366ef4ac7afd4b6cd7e2d94ac35d4f6783701796e5f278c36c51d42c3", "canonical_id": "xcsh-docs:data-sources:site:reference", "child_ids": ["xcsh-docs:data-sources:site:properties:admin_user_credentials", "xcsh-docs:data-sources:site:properties:connected_re", "xcsh-docs:data-sources:site:properties:connected_re_for_config", "xcsh-docs:data-sources:site:properties:coordinates", "xcsh-docs:data-sources:site:properties:default_underlay_network", "xcsh-docs:data-sources:site:properties:kubernetes_upgrade_drain", "xcsh-docs:data-sources:site:properties:main_nodes", "xcsh-docs:data-sources:site:properties:private_connectivity", "xcsh-docs:data-sources:site:properties:re_select", "xcsh-docs:data-sources:site:properties:vip_params_per_az"], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:reference", "parent_id": "xcsh-docs:data-sources:site:fundamentals", "path": "docs/guides/data-sources--site--reference.md", "provider_name": "site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site](../data-sources/site.md)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Computed.

Site's geographical address that can be used to determine its latitude and longitude.

- [admin_user_credentials](data-sources--site--properties--admin_user_credentials.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

<a id="schema-bgp_peer_address"></a>

### bgp_peer_address property

Type: `"string"`. Computed.

Optional BGP peer address that can be used as parameter for BGP configuration when BGP is configured
to fetch BGP peer address from site Object. This can be used to change peer address per site in
fleet.

<a id="schema-bgp_router_id"></a>

### bgp_router_id property

Type: `"string"`. Computed.

Optional BGP router ID that can be used as parameter for BGP configuration when BGP is configured to
fetch BGP router ID from site object. This can be used to change router ID per site in a fleet.

<a id="schema-ce_site_mode"></a>

### ce_site_mode property

Type: `"string"`. Computed.

\[Enum:
CE\_SITE\_MODE\_INGRESS\_EGRESS\_GW|CE\_SITE\_MODE\_INGRESS\_GW|CE\_SITE\_MODE\_EGRESS\_GW|CE\_SITE\_MODE\_DC\_CLOUD\_GW|CE\_SITE\_MODE\_CPE\]
If Site is CE, it can be in following modes Ingress Egress Gateway CE Ingress Gateway CE Egress
Gateway CE DC Cloud Gateway CE CPE CE. Possible values are \`CE\_SITE\_MODE\_INGRESS\_EGRESS\_GW\`,
\`CE\_SITE\_MODE\_INGRESS\_GW\`, \`CE\_SITE\_MODE\_EGRESS\_GW\`, \`CE\_SITE\_MODE\_DC\_CLOUD\_GW\`,
\`CE\_SITE\_MODE\_CPE\`. Defaults to \`CE\_SITE\_MODE\_INGRESS\_EGRESS\_GW\`.

- [connected_re](data-sources--site--properties--connected_re.md): complete subsection reference.

- [connected_re_for_config](data-sources--site--properties--connected_re_for_config.md): complete subsection reference.

- [coordinates](data-sources--site--properties--coordinates.md): complete subsection reference.

- [default_underlay_network](data-sources--site--properties--default_underlay_network.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

<a id="schema-desired_pool_count"></a>

### desired_pool_count property

Type: `"number"`. Computed.

Desired pool count represent desired number of worker(non master) nodes for manual scaling of public
cloud(AWS, GCP, Azure) sites. The desired count must be less than or equal to the maximum size of
the scaling group for a given public cloud. One may also have to increase maximum scaling group..

<a id="schema-global_access_k8s_enabled"></a>

### global_access_k8s_enabled property

Type: `"bool"`. Computed.

Enable or disable functionality flag

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-inside_nameserver"></a>

### inside_nameserver property

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution in inside network.

<a id="schema-inside_vip"></a>

### inside_vip property

Type: `"string"`. Computed.

Optional Virtual IP to be used as automatic VIP for site local inside network. See documentation for
'VIP' in advertise policy to see when Inside VIP is used. When configured, this is used as VIP
(depending on advertise policy configuration).

<a id="schema-ipsec_ssl_nodes_fqdn"></a>

### ipsec_ssl_nodes_fqdn property

Type: `["list", "string"]`. Computed.

FQDN resolves to responders node IP, if there are multiple nodes at site the resolution will give a
list of all/some individual node IP. Multiple FQDN for same site is also allowed.

- [kubernetes_upgrade_drain](data-sources--site--properties--kubernetes_upgrade_drain.md): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-local_access_k8s_enabled"></a>

### local_access_k8s_enabled property

Type: `"bool"`. Computed.

Enable or disable functionality flag

<a id="schema-local_k8s_access_enabled"></a>

### local_k8s_access_enabled property

Type: `"bool"`. Computed.

Lets user know if this site has local K8s cluster enabled via fleet configuration.

- [main_nodes](data-sources--site--properties--main_nodes.md): complete subsection reference.

<a id="schema-multus_enabled"></a>

### multus_enabled property

Type: `"bool"`. Computed.

Indicates that Multus cni is enabled on the site.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Site to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the Site.

<a id="schema-operating_system_version"></a>

### operating_system_version property

Type: `"string"`. Computed.

Desired Operating System version for this site.

<a id="schema-outside_nameserver"></a>

### outside_nameserver property

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution in outside network.

<a id="schema-outside_vip"></a>

### outside_vip property

Type: `"string"`. Computed.

Optional Virtual IP to be used as automatic VIP for site local outside network. See documentation
for 'VIP' in advertise policy to see when Outside VIP is used. When configured, this is used as VIP
(depending on advertise policy configuration).

- [private_connectivity](data-sources--site--properties--private_connectivity.md): complete subsection reference.

- [re_select](data-sources--site--properties--re_select.md): complete subsection reference.

<a id="schema-region"></a>

### region property

Type: `"string"`. Computed.

Cloud Region. A region is a set of datacenters deployed within a latency-defined perimeter and
connected through a dedicated regional low-latency network.

<a id="schema-site_state"></a>

### site_state property

Type: `"string"`. Computed.

\[Enum:
ONLINE|PROVISIONING|UPGRADING|STANDBY|FAILED|REREGISTRATION|WAITINGNODES|DECOMMISSIONING|WAITING\_FOR\_REGISTRATION|ORCHESTRATION\_IN\_PROGRESS|ORCHESTRATION\_COMPLETE|ERROR\_IN\_ORCHESTRATION|DELETING\_CLOUD\_RESOURCES|DELETED\_CLOUD\_RESOURCES|ERROR\_DELETING\_CLOUD\_RESOURCES|VALIDATION\_IN\_PROGRESS|VALIDATION\_SUCCESS|VALIDATION\_FAILED|FAILED\_INACTIVE|UPDATING\_CLOUD\_RESOURCES|ERROR\_UPDATING\_CLOUD\_RESOURCES|ORCHESTRATION\_QUEUED|UPDATE\_QUEUED|DELETE\_QUEUED\]
State of Site defines in which operational state site itself is. Site is online and operational.
Site is in provisioning state. For instance during site deployment or switching to different
connected Regional Edge. Site is in process of upgrade. Possible values are \`ONLINE\`,
\`PROVISIONING\`, \`UPGRADING\`, \`STANDBY\`, \`FAILED\`, \`REREGISTRATION\`, \`WAITINGNODES\`,
\`DECOMMISSIONING\`, \`WAITING\_FOR\_REGISTRATION\`, \`ORCHESTRATION\_IN\_PROGRESS\`,
\`ORCHESTRATION\_COMPLETE\`, \`ERROR\_IN\_ORCHESTRATION\`, \`DELETING\_CLOUD\_RESOURCES\`,
\`DELETED\_CLOUD\_RESOURCES\`, \`ERROR\_DELETING\_CLOUD\_RESOURCES\`, \`VALIDATION\_IN\_PROGRESS\`,
\`VALIDATION\_SUCCESS\`, \`VALIDATION\_FAILED\`, \`FAILED\_INACTIVE\`,
\`UPDATING\_CLOUD\_RESOURCES\`, \`ERROR\_UPDATING\_CLOUD\_RESOURCES\`, \`ORCHESTRATION\_QUEUED\`,
\`UPDATE\_QUEUED\`, \`DELETE\_QUEUED\`. Defaults to \`ONLINE\`.

<a id="schema-site_subtype"></a>

### site_subtype property

Type: `"string"`. Computed.

\[Enum: NO\_SUBTYPE|VES\_IO\_USE\_RE|VES\_IO\_CE\_IN\_K8S|VES\_IO\_HIDDEN\_RE\] Sit Subtype No
Subtype Regional Edge isn't ready yet. Configuration isn't propagated for this site. Regional Edge
which is hidden from customer. Configuration will be propagated. CE running in Kubernetes. Possible
values are \`NO\_SUBTYPE\`, \`VES\_IO\_USE\_RE\`, \`VES\_IO\_CE\_IN\_K8S\`, \`VES\_IO\_HIDDEN\_RE\`.
Defaults to \`NO\_SUBTYPE\`.

<a id="schema-site_to_site_network_type"></a>

### site_to_site_network_type property

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

<a id="schema-site_to_site_tunnel_ip"></a>

### site_to_site_tunnel_ip property

Type: `"string"`. Computed.

Optionsl, VIP in the site\_to\_site\_network\_type configured above used for terminating IPsec/SSL
tunnels created with SiteMeshGroup.

<a id="schema-site_type"></a>

### site_type property

Type: `"string"`. Computed.

\[Enum: INVALID|REGIONAL\_EDGE|CUSTOMER\_EDGE|NGINX\_ONE\] Site Type which can either RE or CE
Invalid type of site Regional Edge site Customer Edge site. Possible values are \`INVALID\`,
\`REGIONAL\_EDGE\`, \`CUSTOMER\_EDGE\`, \`NGINX\_ONE\`.

<a id="schema-tunnel_dead_timeout"></a>

### tunnel_dead_timeout property

Type: `"number"`. Computed.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

<a id="schema-tunnel_type"></a>

### tunnel_type property

Type: `"string"`. Computed.

\[Enum:
SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL|SITE\_TO\_SITE\_TUNNEL\_IPSEC|SITE\_TO\_SITE\_TUNNEL\_SSL\]
Tunnel encapsulation to be used between sites Tunnel can operate in both IPsec and SSL, with IPsec
being preferred over SSL. Tunnel is of type IPsec Tunnel is of type SSL. Possible values are
\`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`, \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\`,
\`SITE\_TO\_SITE\_TUNNEL\_SSL\`. Defaults to \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`.

- [vip_params_per_az](data-sources--site--properties--vip_params_per_az.md): complete subsection reference.

<a id="schema-vip_vrrp_mode"></a>

### vip_vrrp_mode property

Type: `"string"`. Computed.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

<a id="schema-vm_enabled"></a>

### vm_enabled property

Type: `"bool"`. Computed.

Indicates that virtual machine support is enabled on the site.

<a id="schema-volterra_software_override"></a>

### volterra_software_override property

Type: `"string"`. Computed.

\[Enum:
SITE\_SOFTWARE\_OVERRIDE\_SITE|SITE\_SOFTWARE\_OVERRIDE\_NEWER|SITE\_SOFTWARE\_OVERRIDE\_FLEET\]
Decide which software version takes effect in case of conflict between site and fleet Software
version in site will take precedence. Between site and fleet newer software version will take
precedence. Software version in fleet will take precedence. Possible values are
\`SITE\_SOFTWARE\_OVERRIDE\_SITE\`, \`SITE\_SOFTWARE\_OVERRIDE\_NEWER\`,
\`SITE\_SOFTWARE\_OVERRIDE\_FLEET\`. Defaults to \`SITE\_SOFTWARE\_OVERRIDE\_SITE\`.

<a id="schema-volterra_software_version"></a>

### volterra_software_version property

Type: `"string"`. Computed.

Desired F5XC software version for this site, a string matching released set of software components.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--site--reference.md#schema-address) |
| `admin_user_credentials` | [admin_user_credentials](data-sources--site--properties--admin_user_credentials.md#section) |
| `admin_user_credentials.admin_password` | [admin_user_credentials.admin_password](data-sources--site--properties--admin_user_credentials--admin_password.md#section) |
| `admin_user_credentials.admin_password.blindfold_secret_info` | [admin_user_credentials.admin_password.blindfold_secret_info](data-sources--site--properties--admin_user_credentials--admin_password--blindfold_secret_info.md#section) |
| `admin_user_credentials.admin_password.blindfold_secret_info.decryption_provider` | [admin_user_credentials.admin_password.blindfold_secret_info.decryption_provider](data-sources--site--properties--admin_user_credentials--admin_password--blindfold_secret_info.md#schema-admin_user_credentials--admin_password--blindfold_secret_info--decryption_provider) |
| `admin_user_credentials.admin_password.blindfold_secret_info.location` | [admin_user_credentials.admin_password.blindfold_secret_info.location](data-sources--site--properties--admin_user_credentials--admin_password--blindfold_secret_info.md#schema-admin_user_credentials--admin_password--blindfold_secret_info--location) |
| `admin_user_credentials.admin_password.blindfold_secret_info.store_provider` | [admin_user_credentials.admin_password.blindfold_secret_info.store_provider](data-sources--site--properties--admin_user_credentials--admin_password--blindfold_secret_info.md#schema-admin_user_credentials--admin_password--blindfold_secret_info--store_provider) |
| `admin_user_credentials.admin_password.clear_secret_info` | [admin_user_credentials.admin_password.clear_secret_info](data-sources--site--properties--admin_user_credentials--admin_password--clear_secret_info.md#section) |
| `admin_user_credentials.admin_password.clear_secret_info.provider_ref` | [admin_user_credentials.admin_password.clear_secret_info.provider_ref](data-sources--site--properties--admin_user_credentials--admin_password--clear_secret_info.md#schema-admin_user_credentials--admin_password--clear_secret_info--provider_ref) |
| `admin_user_credentials.admin_password.clear_secret_info.url` | [admin_user_credentials.admin_password.clear_secret_info.url](data-sources--site--properties--admin_user_credentials--admin_password--clear_secret_info.md#schema-admin_user_credentials--admin_password--clear_secret_info--url) |
| `admin_user_credentials.ssh_key` | [admin_user_credentials.ssh_key](data-sources--site--properties--admin_user_credentials.md#schema-admin_user_credentials--ssh_key) |
| `annotations` | [annotations](data-sources--site--reference.md#schema-annotations) |
| `bgp_peer_address` | [bgp_peer_address](data-sources--site--reference.md#schema-bgp_peer_address) |
| `bgp_router_id` | [bgp_router_id](data-sources--site--reference.md#schema-bgp_router_id) |
| `ce_site_mode` | [ce_site_mode](data-sources--site--reference.md#schema-ce_site_mode) |
| `connected_re` | [connected_re](data-sources--site--properties--connected_re.md#section) |
| `connected_re.kind` | [connected_re.kind](data-sources--site--properties--connected_re.md#schema-connected_re--kind) |
| `connected_re.name` | [connected_re.name](data-sources--site--properties--connected_re.md#schema-connected_re--name) |
| `connected_re.namespace` | [connected_re.namespace](data-sources--site--properties--connected_re.md#schema-connected_re--namespace) |
| `connected_re.tenant` | [connected_re.tenant](data-sources--site--properties--connected_re.md#schema-connected_re--tenant) |
| `connected_re.uid` | [connected_re.uid](data-sources--site--properties--connected_re.md#schema-connected_re--uid) |
| `connected_re_for_config` | [connected_re_for_config](data-sources--site--properties--connected_re_for_config.md#section) |
| `connected_re_for_config.kind` | [connected_re_for_config.kind](data-sources--site--properties--connected_re_for_config.md#schema-connected_re_for_config--kind) |
| `connected_re_for_config.name` | [connected_re_for_config.name](data-sources--site--properties--connected_re_for_config.md#schema-connected_re_for_config--name) |
| `connected_re_for_config.namespace` | [connected_re_for_config.namespace](data-sources--site--properties--connected_re_for_config.md#schema-connected_re_for_config--namespace) |
| `connected_re_for_config.tenant` | [connected_re_for_config.tenant](data-sources--site--properties--connected_re_for_config.md#schema-connected_re_for_config--tenant) |
| `connected_re_for_config.uid` | [connected_re_for_config.uid](data-sources--site--properties--connected_re_for_config.md#schema-connected_re_for_config--uid) |
| `coordinates` | [coordinates](data-sources--site--properties--coordinates.md#section) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--site--properties--coordinates.md#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--site--properties--coordinates.md#schema-coordinates--longitude) |
| `default_underlay_network` | [default_underlay_network](data-sources--site--properties--default_underlay_network.md#section) |
| `default_underlay_network.site_local_inside` | [default_underlay_network.site_local_inside](data-sources--site--properties--default_underlay_network--site_local_inside.md#section) |
| `default_underlay_network.site_local_outside` | [default_underlay_network.site_local_outside](data-sources--site--properties--default_underlay_network--site_local_outside.md#section) |
| `description` | [description](data-sources--site--reference.md#schema-description) |
| `desired_pool_count` | [desired_pool_count](data-sources--site--reference.md#schema-desired_pool_count) |
| `global_access_k8s_enabled` | [global_access_k8s_enabled](data-sources--site--reference.md#schema-global_access_k8s_enabled) |
| `id` | [id](data-sources--site--reference.md#schema-id) |
| `inside_nameserver` | [inside_nameserver](data-sources--site--reference.md#schema-inside_nameserver) |
| `inside_vip` | [inside_vip](data-sources--site--reference.md#schema-inside_vip) |
| `ipsec_ssl_nodes_fqdn` | [ipsec_ssl_nodes_fqdn](data-sources--site--reference.md#schema-ipsec_ssl_nodes_fqdn) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--site--properties--kubernetes_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--disable_vega_upgrade_mode.md#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--site--properties--kubernetes_upgrade_drain--enable_upgrade_drain--enable_vega_upgrade_mode.md#section) |
| `labels` | [labels](data-sources--site--reference.md#schema-labels) |
| `local_access_k8s_enabled` | [local_access_k8s_enabled](data-sources--site--reference.md#schema-local_access_k8s_enabled) |
| `local_k8s_access_enabled` | [local_k8s_access_enabled](data-sources--site--reference.md#schema-local_k8s_access_enabled) |
| `main_nodes` | [main_nodes](data-sources--site--properties--main_nodes.md#section) |
| `main_nodes.name` | [main_nodes.name](data-sources--site--properties--main_nodes.md#schema-main_nodes--name) |
| `main_nodes.sli_address` | [main_nodes.sli_address](data-sources--site--properties--main_nodes.md#schema-main_nodes--sli_address) |
| `main_nodes.slo_address` | [main_nodes.slo_address](data-sources--site--properties--main_nodes.md#schema-main_nodes--slo_address) |
| `multus_enabled` | [multus_enabled](data-sources--site--reference.md#schema-multus_enabled) |
| `name` | [name](data-sources--site--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--site--reference.md#schema-namespace) |
| `operating_system_version` | [operating_system_version](data-sources--site--reference.md#schema-operating_system_version) |
| `outside_nameserver` | [outside_nameserver](data-sources--site--reference.md#schema-outside_nameserver) |
| `outside_vip` | [outside_vip](data-sources--site--reference.md#schema-outside_vip) |
| `private_connectivity` | [private_connectivity](data-sources--site--properties--private_connectivity.md#section) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--site--properties--private_connectivity--cloud_link.md#section) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--name) |
| `private_connectivity.cloud_link.state` | [private_connectivity.cloud_link.state](data-sources--site--properties--private_connectivity--cloud_link.md#schema-private_connectivity--cloud_link--state) |
| `private_connectivity.private_network_name` | [private_connectivity.private_network_name](data-sources--site--properties--private_connectivity.md#schema-private_connectivity--private_network_name) |
| `re_select` | [re_select](data-sources--site--properties--re_select.md#section) |
| `re_select.geo_proximity` | [re_select.geo_proximity](data-sources--site--properties--re_select--geo_proximity.md#section) |
| `re_select.specific_geography` | [re_select.specific_geography](data-sources--site--properties--re_select.md#schema-re_select--specific_geography) |
| `re_select.specific_re` | [re_select.specific_re](data-sources--site--properties--re_select--specific_re.md#section) |
| `re_select.specific_re.backup_re` | [re_select.specific_re.backup_re](data-sources--site--properties--re_select--specific_re.md#schema-re_select--specific_re--backup_re) |
| `re_select.specific_re.primary_re` | [re_select.specific_re.primary_re](data-sources--site--properties--re_select--specific_re.md#schema-re_select--specific_re--primary_re) |
| `region` | [region](data-sources--site--reference.md#schema-region) |
| `site_state` | [site_state](data-sources--site--reference.md#schema-site_state) |
| `site_subtype` | [site_subtype](data-sources--site--reference.md#schema-site_subtype) |
| `site_to_site_network_type` | [site_to_site_network_type](data-sources--site--reference.md#schema-site_to_site_network_type) |
| `site_to_site_tunnel_ip` | [site_to_site_tunnel_ip](data-sources--site--reference.md#schema-site_to_site_tunnel_ip) |
| `site_type` | [site_type](data-sources--site--reference.md#schema-site_type) |
| `tunnel_dead_timeout` | [tunnel_dead_timeout](data-sources--site--reference.md#schema-tunnel_dead_timeout) |
| `tunnel_type` | [tunnel_type](data-sources--site--reference.md#schema-tunnel_type) |
| `vip_params_per_az` | [vip_params_per_az](data-sources--site--properties--vip_params_per_az.md#section) |
| `vip_params_per_az.az_name` | [vip_params_per_az.az_name](data-sources--site--properties--vip_params_per_az.md#schema-vip_params_per_az--az_name) |
| `vip_params_per_az.inside_vip` | [vip_params_per_az.inside_vip](data-sources--site--properties--vip_params_per_az.md#schema-vip_params_per_az--inside_vip) |
| `vip_params_per_az.inside_vip_cname` | [vip_params_per_az.inside_vip_cname](data-sources--site--properties--vip_params_per_az.md#schema-vip_params_per_az--inside_vip_cname) |
| `vip_params_per_az.inside_vip_v6` | [vip_params_per_az.inside_vip_v6](data-sources--site--properties--vip_params_per_az.md#schema-vip_params_per_az--inside_vip_v6) |
| `vip_params_per_az.outside_vip` | [vip_params_per_az.outside_vip](data-sources--site--properties--vip_params_per_az.md#schema-vip_params_per_az--outside_vip) |
| `vip_params_per_az.outside_vip_cname` | [vip_params_per_az.outside_vip_cname](data-sources--site--properties--vip_params_per_az.md#schema-vip_params_per_az--outside_vip_cname) |
| `vip_params_per_az.outside_vip_v6` | [vip_params_per_az.outside_vip_v6](data-sources--site--properties--vip_params_per_az.md#schema-vip_params_per_az--outside_vip_v6) |
| `vip_vrrp_mode` | [vip_vrrp_mode](data-sources--site--reference.md#schema-vip_vrrp_mode) |
| `vm_enabled` | [vm_enabled](data-sources--site--reference.md#schema-vm_enabled) |
| `volterra_software_override` | [volterra_software_override](data-sources--site--reference.md#schema-volterra_software_override) |
| `volterra_software_version` | [volterra_software_version](data-sources--site--reference.md#schema-volterra_software_version) |

## Next pages

- [admin_user_credentials](data-sources--site--properties--admin_user_credentials.md)
- [connected_re](data-sources--site--properties--connected_re.md)
- [connected_re_for_config](data-sources--site--properties--connected_re_for_config.md)
- [coordinates](data-sources--site--properties--coordinates.md)
- [default_underlay_network](data-sources--site--properties--default_underlay_network.md)
- [kubernetes_upgrade_drain](data-sources--site--properties--kubernetes_upgrade_drain.md)
- [main_nodes](data-sources--site--properties--main_nodes.md)
- [private_connectivity](data-sources--site--properties--private_connectivity.md)
- [re_select](data-sources--site--properties--re_select.md)
- [vip_params_per_az](data-sources--site--properties--vip_params_per_az.md)
- [xcsh_site](../data-sources/site.md)
