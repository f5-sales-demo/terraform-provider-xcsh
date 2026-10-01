---
page_title: "xcsh_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site reference."
---

# xcsh_site reference

<a id="canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b91a8a0759ec02826d0a0c5483397d4801e2c9dfbe591314f08a97409bd8146d"></a>

## Property reference — Property reference / 0ba6da97e504 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- Property reference

<a id="canonical-730022ba1af59798222811d6f403a7f5cf05e36974cadd2fc699217454e98e9a"></a>

## Direct properties — Property reference / 0ba6da97e504 / 3

<a id="canonical-e1edcdbf0d8e265065623f86f5ffc7d0ac4b426006ea5460b423657198319e4b"></a>

<a id="canonical-306ea2be33e547ea5c8a34fea4778308e5b807e262722c17811f5bed26a5e97d"></a>

## address property — Property reference / 0ba6da97e504 / 4

Type: `"string"`. Computed.

Site's geographical address that can be used to determine its latitude and longitude.

- [admin_user_credentials](data-sources--site--reference--group-001.md#canonical-23cada312bda320b7efc68a617dd2908f73a2b8261f34adfc4988ae22d9199ce): complete subsection reference.

<a id="canonical-dc0f3d6568afbc546327f69ebc7aab839521db5a566ab77909efdee846c6e426"></a>

<a id="canonical-7aa57f9963d96bfad35d2c566f7a2c1d4e20c3fa59c9b8725990950d385cf3ff"></a>

## annotations property — Property reference / 0ba6da97e504 / 5

Type: `["map", "string"]`. Computed.

Annotations.

<a id="canonical-a1189e7b3d5d87ce0774e7772c9424e9c26caa09523f393015324d9ba9f01db3"></a>

<a id="canonical-036705e69c175495326e167b5fb25c5f05b3f382a5f8692ddc25314bfc9a75be"></a>

## bgp_peer_address property — Property reference / 0ba6da97e504 / 6

Type: `"string"`. Computed.

Optional BGP peer address that can be used as parameter for BGP configuration when BGP is configured
to fetch BGP peer address from site Object. This can be used to change peer address per site in
fleet.

<a id="canonical-a25f5c93d8e980f4d4c967adb74a1ec317146e35fbb6b2ec61e35d81c2fecfae"></a>

<a id="canonical-f9d4035e931870b55ec0b1e4315d4f393c2d8032a3488c585bc0cd9c1f10cbed"></a>

## bgp_router_id property — Property reference / 0ba6da97e504 / 7

Type: `"string"`. Computed.

Optional BGP router ID that can be used as parameter for BGP configuration when BGP is configured to
fetch BGP router ID from site object. This can be used to change router ID per site in a fleet.

<a id="canonical-19981d07792ba1c659d6bec32bfe2fa0c9ee548db3f42cb47c4d7071e2590f62"></a>

<a id="canonical-34bbbf24a24c82e66d529cde45abda67f2609ff1d450e6c68fca955ef71c2d2a"></a>

## ce_site_mode property — Property reference / 0ba6da97e504 / 8

Type: `"string"`. Computed.

\[Enum:
CE\_SITE\_MODE\_INGRESS\_EGRESS\_GW|CE\_SITE\_MODE\_INGRESS\_GW|CE\_SITE\_MODE\_EGRESS\_GW|CE\_SITE\_MODE\_DC\_CLOUD\_GW|CE\_SITE\_MODE\_CPE\]
If Site is CE, it can be in following modes Ingress Egress Gateway CE Ingress Gateway CE Egress
Gateway CE DC Cloud Gateway CE CPE CE. Possible values are \`CE\_SITE\_MODE\_INGRESS\_EGRESS\_GW\`,
\`CE\_SITE\_MODE\_INGRESS\_GW\`, \`CE\_SITE\_MODE\_EGRESS\_GW\`, \`CE\_SITE\_MODE\_DC\_CLOUD\_GW\`,
\`CE\_SITE\_MODE\_CPE\`. Defaults to \`CE\_SITE\_MODE\_INGRESS\_EGRESS\_GW\`.

- [connected_re](data-sources--site--reference--group-001.md#canonical-85706de1aed9b2978a14e030cedde0be433104392d6a8841ea76bb2dfbeaa589): complete subsection reference.

- [connected_re_for_config](data-sources--site--reference--group-001.md#canonical-70617853cc03e32ab5d93d5623a6b5794c6f8b550e5f402d8a1445c8bba394a4): complete subsection reference.

- [coordinates](data-sources--site--reference--group-001.md#canonical-4acdcf69b09e0ea8e1684d40f9537429ddefa1d373e5328e2ea7fc7e4f2f9faf): complete subsection reference.

- [default_underlay_network](data-sources--site--reference--group-001.md#canonical-bd17269542aad6bf8548924856fd769c4a855887d702aba855aa7688e63cf3f5): complete subsection reference.

<a id="canonical-1aa0740fe289e6156b2f282987f1a6246707c082298b0c774cd1854c599498c2"></a>

<a id="canonical-a2b0deadf8799206d97c1ccf5f1ed880c4c0007de37d93a65b4d844c5eeeb7ca"></a>

## description property — Property reference / 0ba6da97e504 / 9

Type: `"string"`. Computed.

Description.

<a id="canonical-4701dea9755d7fe85de1c270abf24ae95b80cf163833ffb93a880047d3bf799d"></a>

<a id="canonical-9a2bf0caa8ac5b4095375ee3179585865c8f0dc59dfd8fae3b0c87020cd1f2ac"></a>

## desired_pool_count property — Property reference / 0ba6da97e504 / 10

Type: `"number"`. Computed.

Desired pool count represent desired number of worker(non master) nodes for manual scaling of public
cloud(AWS, GCP, Azure) sites. The desired count must be less than or equal to the maximum size of
the scaling group for a given public cloud. One may also have to increase maximum scaling group..

<a id="canonical-82b2541194a96edfb165a32a9d30c96aed7874085195254a19879db8f0948b8e"></a>

<a id="canonical-729171f833e8bdccdcbadd3d99129b189bd85a4d746e67890ce7c24b6cc25289"></a>

## global_access_k8s_enabled property — Property reference / 0ba6da97e504 / 11

Type: `"bool"`. Computed.

Enable or disable functionality flag

<a id="canonical-6014e9e7b0c1ead19e0db77ea656d80dd997fffc494fca3ce3c928b055307cd9"></a>

<a id="canonical-33a5b89c05e8c6bd34613d951064d2248f188092280225ebf30b680fe07969b4"></a>

## id property — Property reference / 0ba6da97e504 / 12

Type: `"string"`. Computed.

Unique identifier.

<a id="canonical-835345b0849e849a4fca0028e140f836d025ddaa2439243203b4ab0ef49207c2"></a>

<a id="canonical-2c2911d674b3c44ac54ddee358d713e00ba1c958477d3b3e24289e1af2af983e"></a>

## inside_nameserver property — Property reference / 0ba6da97e504 / 13

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution in inside network.

<a id="canonical-48127a41799ba8399ce5b39dd28f5814b7086fe1a340e21988e48491a45f5254"></a>

<a id="canonical-ae56f5f85d7053779e0b5885abfa0fbd7618b38a3c4c9baa9704ca9024dc5ddb"></a>

## inside_vip property — Property reference / 0ba6da97e504 / 14

Type: `"string"`. Computed.

Optional Virtual IP to be used as automatic VIP for site local inside network. See documentation for
'VIP' in advertise policy to see when Inside VIP is used. When configured, this is used as VIP
(depending on advertise policy configuration).

<a id="canonical-3520fb34b7065f08d4ffd16c39d742d83b4e731ca8e7a153b5c4c07fea78f724"></a>

<a id="canonical-9d8e4d2f84e82f38bc23a028c405d8d730d8e044280f70caa8f60dddc19e47a3"></a>

## ipsec_ssl_nodes_fqdn property — Property reference / 0ba6da97e504 / 15

Type: `["list", "string"]`. Computed.

FQDN resolves to responders node IP, if there are multiple nodes at site the resolution will give a
list of all/some individual node IP. Multiple FQDN for same site is also allowed.

- [kubernetes_upgrade_drain](data-sources--site--reference--group-001.md#canonical-82ce9fbbef5339fa3591881ed6c4c9c12363dc591088c814863b0e183f8ad380): complete subsection reference.

<a id="canonical-b1b8e7293e1db75693f346144a4a8029ebd9f3884b8cff90d2f49bfc32ca5c23"></a>

<a id="canonical-f66c62de446b144b18d4368bd4fc4ce014bfd9ccec41ba4213ec843de36b29f7"></a>

## labels property — Property reference / 0ba6da97e504 / 16

Type: `["map", "string"]`. Computed.

Labels.

<a id="canonical-1a41503c77158a53a9f2b3fbce96d62a7038e6232fcf263f991f16bb1fe3e2ee"></a>

<a id="canonical-cef9baafbc1098d4858b7df280b7111301ba16d846c8b2ba6f4ea889f65e7dff"></a>

## local_access_k8s_enabled property — Property reference / 0ba6da97e504 / 17

Type: `"bool"`. Computed.

Enable or disable functionality flag

<a id="canonical-e9309e289d7405c9e1feb14ace890f6ba6d00dea8387079db677ca2d214ad8ab"></a>

<a id="canonical-29e66ab56214e8ab84e661a4d212b4734111f3e39da56cd2b746bc16c7f1aad1"></a>

## local_k8s_access_enabled property — Property reference / 0ba6da97e504 / 18

Type: `"bool"`. Computed.

Lets user know if this site has local K8s cluster enabled via fleet configuration.

- [main_nodes](data-sources--site--reference--group-001.md#canonical-e659811d82fb9727fa18dff418f5ce4c62d4b56265159514aabefeaf9e7885da): complete subsection reference.

<a id="canonical-c226eeb4afcede522d549d37b6550101c32e498f21aed605aa213dd09e4648b1"></a>

<a id="canonical-9131cab63335ab28a8afd22acea5cac34eda1c551de0f16d56a6395409aa5ce9"></a>

## multus_enabled property — Property reference / 0ba6da97e504 / 19

Type: `"bool"`. Computed.

Indicates that Multus cni is enabled on the site.

<a id="canonical-0e081d1ccafdd9419b42618cc8fdb922b3f5ef98bdd1dacf1a6294bd271a75e0"></a>

<a id="canonical-e934758c72d28998a041198cbbba8516c4caf1623aed57f4c193e4bf2f3b65b2"></a>

## name property — Property reference / 0ba6da97e504 / 20

Type: `"string"`. Required.

Name of the Site to look up.

<a id="canonical-af9323d78a9e8ffc49c410f7893d78eb6640d89165bb8896b720e2bb0485f404"></a>

<a id="canonical-f82349cf6ba50498f9c9a4554efaa11d24a10c043f9774ab3c07553bff184b54"></a>

## namespace property — Property reference / 0ba6da97e504 / 21

Type: `"string"`. Required.

Namespace of the Site.

<a id="canonical-2adadabe99beccdf90de02352f631fa9f35524165e965077cea2b0c5fe147504"></a>

<a id="canonical-f76def806c7bbee79d08c9943e34333d11da1e721bb41f87d2bd8f8c4672d6c6"></a>

## operating_system_version property — Property reference / 0ba6da97e504 / 22

Type: `"string"`. Computed.

Desired Operating System version for this site.

<a id="canonical-68f168a53088bc8028e164f7ab3daed028e5dcf545dd6c4ea3cb432c58501fae"></a>

<a id="canonical-839024c7bae45d4a6cb512484b4bf5d69b28b977abfc471b35269b8d8548218f"></a>

## outside_nameserver property — Property reference / 0ba6da97e504 / 23

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution in outside network.

<a id="canonical-221149bbf67ce805773daa5f9a85987564daa53ed7eb04ea21b573204cb1f906"></a>

<a id="canonical-56dc39521c5a620f4d892b3cb1609e6724f49c40f01f9127ecfcf082ca72c6ed"></a>

## outside_vip property — Property reference / 0ba6da97e504 / 24

Type: `"string"`. Computed.

Optional Virtual IP to be used as automatic VIP for site local outside network. See documentation
for 'VIP' in advertise policy to see when Outside VIP is used. When configured, this is used as VIP
(depending on advertise policy configuration).

- [private_connectivity](data-sources--site--reference--group-001.md#canonical-c156e10b94775d615e04b930cd431c283b2f53cf610d0eb11028b58aa65a8c55): complete subsection reference.

- [re_select](data-sources--site--reference--group-001.md#canonical-b1e08c929b5c34e6b18d034595249b5dd6e81bb071980608b29fbe99c33a6fcf): complete subsection reference.

<a id="canonical-0335e124f3d5cce1dc0ea4b921d8663f75562d0b8870e24ec3782f0fe6c6f6e9"></a>

<a id="canonical-ce0a2f6f894ad51dd500aec29fccb6389cc213fbd9921dfeaed5f83cba3591e8"></a>

## region property — Property reference / 0ba6da97e504 / 25

Type: `"string"`. Computed.

Cloud Region. A region is a set of datacenters deployed within a latency-defined perimeter and
connected through a dedicated regional low-latency network.

<a id="canonical-02a95bca7833ac30b241e3a1bad210e9d6f11c6886a0ea95d50cfac8559b3bd0"></a>

<a id="canonical-fb9010b5b7962aa586605450d2f35c82f06c878f1f05c9df661d243608a6307a"></a>

## site_state property — Property reference / 0ba6da97e504 / 26

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

<a id="canonical-6167a55281b1adebaa5321f389f9703f9d67e09ebb1bdd8480e194145c732508"></a>

<a id="canonical-54f5a71a4b3ba99dd4585a2890b6bd7b85d613b70542388d76be48431af4dfc7"></a>

## site_subtype property — Property reference / 0ba6da97e504 / 27

Type: `"string"`. Computed.

\[Enum: NO\_SUBTYPE|VES\_IO\_USE\_RE|VES\_IO\_CE\_IN\_K8S|VES\_IO\_HIDDEN\_RE\] Sit Subtype No
Subtype Regional Edge isn't ready yet. Configuration isn't propagated for this site. Regional Edge
which is hidden from customer. Configuration will be propagated. CE running in Kubernetes. Possible
values are \`NO\_SUBTYPE\`, \`VES\_IO\_USE\_RE\`, \`VES\_IO\_CE\_IN\_K8S\`, \`VES\_IO\_HIDDEN\_RE\`.
Defaults to \`NO\_SUBTYPE\`.

<a id="canonical-ccb5d5c958daf8c6667deddbe8c3a0dee28c8dd02546a4b6a1a058b92e7c8c16"></a>

<a id="canonical-5116deaad5a6ac64009328bff82b0eeb2823f075a14c53ef9a4f0bc11525d518"></a>

## site_to_site_network_type property — Property reference / 0ba6da97e504 / 28

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

<a id="canonical-d3962261b2ea5474eaffc7140dad2425e5cc31053bad59d32931107256c06df1"></a>

<a id="canonical-ecf1cfe558124822b013706515b6f787e3d2d1a9d6649ba3f5554205d322fc5e"></a>

## site_to_site_tunnel_ip property — Property reference / 0ba6da97e504 / 29

Type: `"string"`. Computed.

Optionsl, VIP in the site\_to\_site\_network\_type configured above used for terminating IPsec/SSL
tunnels created with SiteMeshGroup.

<a id="canonical-865c7abe4173bd3309f9d4e91758ce01b30822cdc4b8385184db4254f080ecb4"></a>

<a id="canonical-e4c7d8b3bc06fcd5d72efa2b34f84b5f5f1d26f5dd56c0610b7e6baabfd6fd82"></a>

## site_type property — Property reference / 0ba6da97e504 / 30

Type: `"string"`. Computed.

\[Enum: INVALID|REGIONAL\_EDGE|CUSTOMER\_EDGE|NGINX\_ONE\] Site Type which can either RE or CE
Invalid type of site Regional Edge site Customer Edge site. Possible values are \`INVALID\`,
\`REGIONAL\_EDGE\`, \`CUSTOMER\_EDGE\`, \`NGINX\_ONE\`.

<a id="canonical-da3a4a09c93fd8c6b3e16ad8174a7b5c3bda5723f283773178b4b5e3ea0df4c9"></a>

<a id="canonical-94301f5539f572b29249b9e9696fd66a9f440566696813058492fcc457d9ea99"></a>

## tunnel_dead_timeout property — Property reference / 0ba6da97e504 / 31

Type: `"number"`. Computed.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

<a id="canonical-c268d80ccfbe930dc6597915632d7fe89b853fa12ad6f1f6fa8295f33974c1d9"></a>

<a id="canonical-054e57cc023e56ca9ea5d4fa511412db45136048afe2f6d4ba15372f71d40814"></a>

## tunnel_type property — Property reference / 0ba6da97e504 / 32

Type: `"string"`. Computed.

\[Enum:
SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL|SITE\_TO\_SITE\_TUNNEL\_IPSEC|SITE\_TO\_SITE\_TUNNEL\_SSL\]
Tunnel encapsulation to be used between sites Tunnel can operate in both IPsec and SSL, with IPsec
being preferred over SSL. Tunnel is of type IPsec Tunnel is of type SSL. Possible values are
\`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`, \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\`,
\`SITE\_TO\_SITE\_TUNNEL\_SSL\`. Defaults to \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`.

- [vip_params_per_az](data-sources--site--reference--group-001.md#canonical-df84def868cc73d6ffffc568e2956d1cd463d618f8305e6d1e14656bd17948da): complete subsection reference.

<a id="canonical-0da6daac52da255f4ceb0058d240712ab35f2fa0537796da15db2809b78a8528"></a>

<a id="canonical-d9e775ae9b93357357e10af0b43716a0f2e43adad8931c958f4d0a594f5bbfa4"></a>

## vip_vrrp_mode property — Property reference / 0ba6da97e504 / 33

Type: `"string"`. Computed.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

<a id="canonical-87d6d3e2c20227eda681f15270abf438e902b42f95cad6ac9101d2ea51ec70c0"></a>

<a id="canonical-cf2e61e9e281fb4543b948ed55b18ac45a1b9e2a74690a08d46ffd386d6e0db9"></a>

## vm_enabled property — Property reference / 0ba6da97e504 / 34

Type: `"bool"`. Computed.

Indicates that virtual machine support is enabled on the site.

<a id="canonical-7952f9b08dceaec3194912b0c374b86877f72f5e1ce0c461e17ad52cceb5aefe"></a>

<a id="canonical-69ae41aa8687f8db70827d24c0b80f0c43f02d1475c36cd805da3d9cc99fd881"></a>

## volterra_software_override property — Property reference / 0ba6da97e504 / 35

Type: `"string"`. Computed.

\[Enum:
SITE\_SOFTWARE\_OVERRIDE\_SITE|SITE\_SOFTWARE\_OVERRIDE\_NEWER|SITE\_SOFTWARE\_OVERRIDE\_FLEET\]
Decide which software version takes effect in case of conflict between site and fleet Software
version in site will take precedence. Between site and fleet newer software version will take
precedence. Software version in fleet will take precedence. Possible values are
\`SITE\_SOFTWARE\_OVERRIDE\_SITE\`, \`SITE\_SOFTWARE\_OVERRIDE\_NEWER\`,
\`SITE\_SOFTWARE\_OVERRIDE\_FLEET\`. Defaults to \`SITE\_SOFTWARE\_OVERRIDE\_SITE\`.

<a id="canonical-8fc63d7b179ac1e136117eff334385f9ecc23bab19b1c3f248e7067ca5de601f"></a>

<a id="canonical-1783a0d4c4c8d43956304c84eee1282d2ab2eba196022ceb71fdc1077de1eb64"></a>

## volterra_software_version property — Property reference / 0ba6da97e504 / 36

Type: `"string"`. Computed.

Desired F5XC software version for this site, a string matching released set of software components.

<a id="canonical-cdb226824a3a2bed0918ed3639bdf2cbd0a9c3fc5c390b83a7d80499591169d9"></a>

## All schema paths — Property reference / 0ba6da97e504 / 37

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--site--reference--group-001.md#canonical-e1edcdbf0d8e265065623f86f5ffc7d0ac4b426006ea5460b423657198319e4b) |
| `admin_user_credentials` | [admin_user_credentials](data-sources--site--reference--group-001.md#canonical-e9ad41bd99ba0698328622f3272de87770af5f50de03f53e1b0b77569e75c599) |
| `admin_user_credentials.admin_password` | [admin_user_credentials.admin_password](data-sources--site--reference--group-001.md#canonical-b8b3bef1e8d1e9cbecc1d7f866dba33a7fa6389210c54aae7a64ba561c040a53) |
| `admin_user_credentials.admin_password.blindfold_secret_info` | [admin_user_credentials.admin_password.blindfold_secret_info](data-sources--site--reference--group-001.md#canonical-88ceed83ab3ef478c3f4830fae173388c2e9598eff9866cb8759c495673b43e0) |
| `admin_user_credentials.admin_password.blindfold_secret_info.decryption_provider` | [admin_user_credentials.admin_password.blindfold_secret_info.decryption_provider](data-sources--site--reference--group-001.md#canonical-9fec5de6cc3a88d55daaccb0a910d15e62440660dd449eeb1f8bcd33b33de212) |
| `admin_user_credentials.admin_password.blindfold_secret_info.location` | [admin_user_credentials.admin_password.blindfold_secret_info.location](data-sources--site--reference--group-001.md#canonical-90b1a0bd73f0f823a2753e959ccfbc3357050b2844372093dcda19d97a3fec1f) |
| `admin_user_credentials.admin_password.blindfold_secret_info.store_provider` | [admin_user_credentials.admin_password.blindfold_secret_info.store_provider](data-sources--site--reference--group-001.md#canonical-3dbc5da04df28793b8be6bb5d80b15157be06159c8bcab4172dfae817a4992e4) |
| `admin_user_credentials.admin_password.clear_secret_info` | [admin_user_credentials.admin_password.clear_secret_info](data-sources--site--reference--group-001.md#canonical-c14b3d6573633b52179abc4ae97c8eec1f6fabc9009487e4365ca473d9dafe00) |
| `admin_user_credentials.admin_password.clear_secret_info.provider_ref` | [admin_user_credentials.admin_password.clear_secret_info.provider_ref](data-sources--site--reference--group-001.md#canonical-2033164fd940b85d87f97e7925f666ff6856a8d724c0fad265c5ad673ffd98db) |
| `admin_user_credentials.admin_password.clear_secret_info.url` | [admin_user_credentials.admin_password.clear_secret_info.url](data-sources--site--reference--group-001.md#canonical-36daa3f0efc1ce8c564a9dfea21e13851ff10ad42515a779206837f0bf860028) |
| `admin_user_credentials.ssh_key` | [admin_user_credentials.ssh_key](data-sources--site--reference--group-001.md#canonical-192d670828ae5d42869a91efb1eb542cb85817330626938523cd1bb6b868b96c) |
| `annotations` | [annotations](data-sources--site--reference--group-001.md#canonical-dc0f3d6568afbc546327f69ebc7aab839521db5a566ab77909efdee846c6e426) |
| `bgp_peer_address` | [bgp_peer_address](data-sources--site--reference--group-001.md#canonical-a1189e7b3d5d87ce0774e7772c9424e9c26caa09523f393015324d9ba9f01db3) |
| `bgp_router_id` | [bgp_router_id](data-sources--site--reference--group-001.md#canonical-a25f5c93d8e980f4d4c967adb74a1ec317146e35fbb6b2ec61e35d81c2fecfae) |
| `ce_site_mode` | [ce_site_mode](data-sources--site--reference--group-001.md#canonical-19981d07792ba1c659d6bec32bfe2fa0c9ee548db3f42cb47c4d7071e2590f62) |
| `connected_re` | [connected_re](data-sources--site--reference--group-001.md#canonical-a7d91e199cdbd33164b316d6de5f60318841d45d797729c40eb6e444800ae63d) |
| `connected_re.kind` | [connected_re.kind](data-sources--site--reference--group-001.md#canonical-4e21f5cc085189e431aa797b96ed0d73a4fe8207ed18def6de422d1be826c3b6) |
| `connected_re.name` | [connected_re.name](data-sources--site--reference--group-001.md#canonical-aff8fca639bddead506db574f7a6ce6220b4732b11cd94457e8e32fd28dc84fc) |
| `connected_re.namespace` | [connected_re.namespace](data-sources--site--reference--group-001.md#canonical-fe615b7fd5875de442912cb7e5a07567bea16475275136dd33bc5ada5520d2e9) |
| `connected_re.tenant` | [connected_re.tenant](data-sources--site--reference--group-001.md#canonical-a8c28fd0d47b7670bcc7a91a2a4ee11afc565887537c04a67d70c210f63a10ca) |
| `connected_re.uid` | [connected_re.uid](data-sources--site--reference--group-001.md#canonical-7b756f810be1acdadc0b89d87eb14c1016477bd2593b1b84387d94d5a9e941c3) |
| `connected_re_for_config` | [connected_re_for_config](data-sources--site--reference--group-001.md#canonical-97421362f29939bb8ef71a2031f1be9bbb20c90f0f51cfa1101f50771490e494) |
| `connected_re_for_config.kind` | [connected_re_for_config.kind](data-sources--site--reference--group-001.md#canonical-f273285434b4b64d3c26f09179625ee704bb055e02bf77387de6afd347aef11d) |
| `connected_re_for_config.name` | [connected_re_for_config.name](data-sources--site--reference--group-001.md#canonical-76e16a31ea9bdc0aae6ea34f9772ee6230bf9cd6d22a95d4ac2e6479eb2dfaee) |
| `connected_re_for_config.namespace` | [connected_re_for_config.namespace](data-sources--site--reference--group-001.md#canonical-9c7ddb44bd40ffa4b50997bc591082f41b16152478f3a82a85672044d3b3fa87) |
| `connected_re_for_config.tenant` | [connected_re_for_config.tenant](data-sources--site--reference--group-001.md#canonical-4daa101ef9a3786128e1b82b3dc91b94442d1420ae9dcda7a653261fb6a22eb4) |
| `connected_re_for_config.uid` | [connected_re_for_config.uid](data-sources--site--reference--group-001.md#canonical-41f6e3f5f5f1a2c2e279508d3ada0b7db91bfa33c50ac1118e404fbd4c0205d4) |
| `coordinates` | [coordinates](data-sources--site--reference--group-001.md#canonical-03b01fed26df694dbf42e985da27d2e75b3f306656a5bdd4816ea426cf76c736) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--site--reference--group-001.md#canonical-0bfe4b6ae704a43dd1ae9dc136d9477ea48e2a846916d4861bfbf9cacebde14c) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--site--reference--group-001.md#canonical-22a44e5606705a32814ee60eb588b4a5884942c98b850bcaea32d9b126a757ae) |
| `default_underlay_network` | [default_underlay_network](data-sources--site--reference--group-001.md#canonical-a762a1a85bca5dbefaab87be71767cb764b9abada86afc9b6176b443cea0e0e8) |
| `default_underlay_network.site_local_inside` | [default_underlay_network.site_local_inside](data-sources--site--reference--group-001.md#canonical-fe41799b0711b493e75b835e956c44a1b6395b60d9eb2ea09fd18d23c23065dc) |
| `default_underlay_network.site_local_outside` | [default_underlay_network.site_local_outside](data-sources--site--reference--group-001.md#canonical-1253b80574300d4d0a3854f83dcd57f7f85d82d342be942f2c2cb258fb98e627) |
| `description` | [description](data-sources--site--reference--group-001.md#canonical-1aa0740fe289e6156b2f282987f1a6246707c082298b0c774cd1854c599498c2) |
| `desired_pool_count` | [desired_pool_count](data-sources--site--reference--group-001.md#canonical-4701dea9755d7fe85de1c270abf24ae95b80cf163833ffb93a880047d3bf799d) |
| `global_access_k8s_enabled` | [global_access_k8s_enabled](data-sources--site--reference--group-001.md#canonical-82b2541194a96edfb165a32a9d30c96aed7874085195254a19879db8f0948b8e) |
| `id` | [id](data-sources--site--reference--group-001.md#canonical-6014e9e7b0c1ead19e0db77ea656d80dd997fffc494fca3ce3c928b055307cd9) |
| `inside_nameserver` | [inside_nameserver](data-sources--site--reference--group-001.md#canonical-835345b0849e849a4fca0028e140f836d025ddaa2439243203b4ab0ef49207c2) |
| `inside_vip` | [inside_vip](data-sources--site--reference--group-001.md#canonical-48127a41799ba8399ce5b39dd28f5814b7086fe1a340e21988e48491a45f5254) |
| `ipsec_ssl_nodes_fqdn` | [ipsec_ssl_nodes_fqdn](data-sources--site--reference--group-001.md#canonical-3520fb34b7065f08d4ffd16c39d742d83b4e731ca8e7a153b5c4c07fea78f724) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--site--reference--group-001.md#canonical-71d124a4030aacd174f374d619778ec41f397171cd6a34b74db40be4f074075c) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-2d608f33760dea286591a976a58d6a23adf9fd3459f458ae8b007c5cdebbde85) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-3a5f0d0538e4764a35eb84d26a4a96346bd8c93beea61011f9e1746868cc114e) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--site--reference--group-001.md#canonical-d02e64fe3052930a946b69e9812d31b4c5a6ae36720436ca682e4a30fc18bb9e) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--site--reference--group-001.md#canonical-bdb3194819542b209c7db0bcad89876c8a83d49de55a959a3d6da57701b9e058) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--site--reference--group-001.md#canonical-84f262c0a2973f3e5431e5c78d5b51725eb2aada3e8427dc26bcfec9b934481b) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--site--reference--group-001.md#canonical-bdfa3c7c4b9c935408578410a4fc3b402eefb161fc46c49f29440218b6b43aa7) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--site--reference--group-001.md#canonical-7f4cd3224ac315d2989cfc1ac3891656dee07d004b5fc2c4997ec9c32cb143f0) |
| `labels` | [labels](data-sources--site--reference--group-001.md#canonical-b1b8e7293e1db75693f346144a4a8029ebd9f3884b8cff90d2f49bfc32ca5c23) |
| `local_access_k8s_enabled` | [local_access_k8s_enabled](data-sources--site--reference--group-001.md#canonical-1a41503c77158a53a9f2b3fbce96d62a7038e6232fcf263f991f16bb1fe3e2ee) |
| `local_k8s_access_enabled` | [local_k8s_access_enabled](data-sources--site--reference--group-001.md#canonical-e9309e289d7405c9e1feb14ace890f6ba6d00dea8387079db677ca2d214ad8ab) |
| `main_nodes` | [main_nodes](data-sources--site--reference--group-001.md#canonical-756f1e1a026215d585567ac0c026d9eed0d8ad0bc9f5c0f73845313eddbc3d14) |
| `main_nodes.name` | [main_nodes.name](data-sources--site--reference--group-001.md#canonical-e3552c5824655418a6ab2bb64e153227f3a6e9857a63370e10f4d0a96f62d5d7) |
| `main_nodes.sli_address` | [main_nodes.sli_address](data-sources--site--reference--group-001.md#canonical-c28e28e21851381098e47ff4a1e201b089dc920b7b14c1246bb62f9e3a1f5967) |
| `main_nodes.slo_address` | [main_nodes.slo_address](data-sources--site--reference--group-001.md#canonical-47bb30fed57cc0f141d5fa281baa4842037b2f074cb937a6cc0b73d226074729) |
| `multus_enabled` | [multus_enabled](data-sources--site--reference--group-001.md#canonical-c226eeb4afcede522d549d37b6550101c32e498f21aed605aa213dd09e4648b1) |
| `name` | [name](data-sources--site--reference--group-001.md#canonical-0e081d1ccafdd9419b42618cc8fdb922b3f5ef98bdd1dacf1a6294bd271a75e0) |
| `namespace` | [namespace](data-sources--site--reference--group-001.md#canonical-af9323d78a9e8ffc49c410f7893d78eb6640d89165bb8896b720e2bb0485f404) |
| `operating_system_version` | [operating_system_version](data-sources--site--reference--group-001.md#canonical-2adadabe99beccdf90de02352f631fa9f35524165e965077cea2b0c5fe147504) |
| `outside_nameserver` | [outside_nameserver](data-sources--site--reference--group-001.md#canonical-68f168a53088bc8028e164f7ab3daed028e5dcf545dd6c4ea3cb432c58501fae) |
| `outside_vip` | [outside_vip](data-sources--site--reference--group-001.md#canonical-221149bbf67ce805773daa5f9a85987564daa53ed7eb04ea21b573204cb1f906) |
| `private_connectivity` | [private_connectivity](data-sources--site--reference--group-001.md#canonical-443c7f09d056cf32a0c1f563d105c01a3bbb2a9676b07da28e5c6c0b4252db63) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--site--reference--group-001.md#canonical-884826ddd9549ed200f72212ce746e3d098b33b597fe7a6f44e3d14149c81ca3) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--site--reference--group-001.md#canonical-16385b40c01df2dd292b41761af77681a9599e8d4e9310c04c85ea14ade42fe5) |
| `private_connectivity.cloud_link.state` | [private_connectivity.cloud_link.state](data-sources--site--reference--group-001.md#canonical-18c57637c245572d111dfc52d7c50378a084ff39f4bd521b3354c0e360832a97) |
| `private_connectivity.private_network_name` | [private_connectivity.private_network_name](data-sources--site--reference--group-001.md#canonical-6bb2f74a9d0f645c0e6987a1352c46becc1351789a9e699524e624dd642f7098) |
| `re_select` | [re_select](data-sources--site--reference--group-001.md#canonical-70af337bc649fc0197600ab7da9df9b1b426de15d2afa4c8212573ca0f8d4d88) |
| `re_select.geo_proximity` | [re_select.geo_proximity](data-sources--site--reference--group-001.md#canonical-66c8c4bb05e8f25cf773729131aa627c4eddb5ff59cb6be18590475af11804c0) |
| `re_select.specific_geography` | [re_select.specific_geography](data-sources--site--reference--group-001.md#canonical-207ca62280fb3c03ef9d3ac13e323a8b56fa4927562805beaa06ec7b7a09a9d7) |
| `re_select.specific_re` | [re_select.specific_re](data-sources--site--reference--group-001.md#canonical-7e8edcd4b37f7b746b9b801b761e6fb6e97f29e9797ad49c3b67f84b94b24a7b) |
| `re_select.specific_re.backup_re` | [re_select.specific_re.backup_re](data-sources--site--reference--group-001.md#canonical-3ccf72c32648e873c90685086acd852d65656fb3fdaee62ecd9339a4108fd037) |
| `re_select.specific_re.primary_re` | [re_select.specific_re.primary_re](data-sources--site--reference--group-001.md#canonical-e008d4f98f96d72fb9c58b96aeee8c7dab95d20bac2055ecf8a823bc3b629be5) |
| `region` | [region](data-sources--site--reference--group-001.md#canonical-0335e124f3d5cce1dc0ea4b921d8663f75562d0b8870e24ec3782f0fe6c6f6e9) |
| `site_state` | [site_state](data-sources--site--reference--group-001.md#canonical-02a95bca7833ac30b241e3a1bad210e9d6f11c6886a0ea95d50cfac8559b3bd0) |
| `site_subtype` | [site_subtype](data-sources--site--reference--group-001.md#canonical-6167a55281b1adebaa5321f389f9703f9d67e09ebb1bdd8480e194145c732508) |
| `site_to_site_network_type` | [site_to_site_network_type](data-sources--site--reference--group-001.md#canonical-ccb5d5c958daf8c6667deddbe8c3a0dee28c8dd02546a4b6a1a058b92e7c8c16) |
| `site_to_site_tunnel_ip` | [site_to_site_tunnel_ip](data-sources--site--reference--group-001.md#canonical-d3962261b2ea5474eaffc7140dad2425e5cc31053bad59d32931107256c06df1) |
| `site_type` | [site_type](data-sources--site--reference--group-001.md#canonical-865c7abe4173bd3309f9d4e91758ce01b30822cdc4b8385184db4254f080ecb4) |
| `tunnel_dead_timeout` | [tunnel_dead_timeout](data-sources--site--reference--group-001.md#canonical-da3a4a09c93fd8c6b3e16ad8174a7b5c3bda5723f283773178b4b5e3ea0df4c9) |
| `tunnel_type` | [tunnel_type](data-sources--site--reference--group-001.md#canonical-c268d80ccfbe930dc6597915632d7fe89b853fa12ad6f1f6fa8295f33974c1d9) |
| `vip_params_per_az` | [vip_params_per_az](data-sources--site--reference--group-001.md#canonical-975f0278fcff1c3020efb711819a9649dbd67a542f1323be97b894196510cc61) |
| `vip_params_per_az.az_name` | [vip_params_per_az.az_name](data-sources--site--reference--group-001.md#canonical-1f6a6d15bcd16477a7ead7934dc1921cae6bf26745769cd0ec5470f2cf07caa9) |
| `vip_params_per_az.inside_vip` | [vip_params_per_az.inside_vip](data-sources--site--reference--group-001.md#canonical-116cafdd64a3ad1a3148440312334f88460ff8f1c845e6434aaea6d95ef69413) |
| `vip_params_per_az.inside_vip_cname` | [vip_params_per_az.inside_vip_cname](data-sources--site--reference--group-001.md#canonical-48b73b24cd3401657f72e99bf5452ffa21824f42bbd8cf4032e8921dbd8740ba) |
| `vip_params_per_az.inside_vip_v6` | [vip_params_per_az.inside_vip_v6](data-sources--site--reference--group-001.md#canonical-5d8ec8c07494b3c2bd4cbe454fddd165a2ae83c75a562e792ce81e7d5e2df967) |
| `vip_params_per_az.outside_vip` | [vip_params_per_az.outside_vip](data-sources--site--reference--group-001.md#canonical-890522a3682f28c289c6a77fd10dd9f835634be8486588b6704dae66cfdf45ca) |
| `vip_params_per_az.outside_vip_cname` | [vip_params_per_az.outside_vip_cname](data-sources--site--reference--group-001.md#canonical-45fb01b0f79b4ae640273c78e22d5c4a98e8657238a5b56547c3e48e2daf6c68) |
| `vip_params_per_az.outside_vip_v6` | [vip_params_per_az.outside_vip_v6](data-sources--site--reference--group-001.md#canonical-ff7bc32c2422d83ce0bf8a126e4341f51e2bbdb1aeb15b516f015fc91e71014d) |
| `vip_vrrp_mode` | [vip_vrrp_mode](data-sources--site--reference--group-001.md#canonical-0da6daac52da255f4ceb0058d240712ab35f2fa0537796da15db2809b78a8528) |
| `vm_enabled` | [vm_enabled](data-sources--site--reference--group-001.md#canonical-87d6d3e2c20227eda681f15270abf438e902b42f95cad6ac9101d2ea51ec70c0) |
| `volterra_software_override` | [volterra_software_override](data-sources--site--reference--group-001.md#canonical-7952f9b08dceaec3194912b0c374b86877f72f5e1ce0c461e17ad52cceb5aefe) |
| `volterra_software_version` | [volterra_software_version](data-sources--site--reference--group-001.md#canonical-8fc63d7b179ac1e136117eff334385f9ecc23bab19b1c3f248e7067ca5de601f) |

<a id="canonical-592b4f29b5be3c9b81336bbddff95a32dfe9a5208be519cbe5205446813fe1ef"></a>

## Next pages — Property reference / 0ba6da97e504 / 38

- [admin_user_credentials](data-sources--site--reference--group-001.md#canonical-23cada312bda320b7efc68a617dd2908f73a2b8261f34adfc4988ae22d9199ce)
- [connected_re](data-sources--site--reference--group-001.md#canonical-85706de1aed9b2978a14e030cedde0be433104392d6a8841ea76bb2dfbeaa589)
- [connected_re_for_config](data-sources--site--reference--group-001.md#canonical-70617853cc03e32ab5d93d5623a6b5794c6f8b550e5f402d8a1445c8bba394a4)
- [coordinates](data-sources--site--reference--group-001.md#canonical-4acdcf69b09e0ea8e1684d40f9537429ddefa1d373e5328e2ea7fc7e4f2f9faf)
- [default_underlay_network](data-sources--site--reference--group-001.md#canonical-bd17269542aad6bf8548924856fd769c4a855887d702aba855aa7688e63cf3f5)
- [kubernetes_upgrade_drain](data-sources--site--reference--group-001.md#canonical-82ce9fbbef5339fa3591881ed6c4c9c12363dc591088c814863b0e183f8ad380)
- [main_nodes](data-sources--site--reference--group-001.md#canonical-e659811d82fb9727fa18dff418f5ce4c62d4b56265159514aabefeaf9e7885da)
- [private_connectivity](data-sources--site--reference--group-001.md#canonical-c156e10b94775d615e04b930cd431c283b2f53cf610d0eb11028b58aa65a8c55)
- [re_select](data-sources--site--reference--group-001.md#canonical-b1e08c929b5c34e6b18d034595249b5dd6e81bb071980608b29fbe99c33a6fcf)
- [vip_params_per_az](data-sources--site--reference--group-001.md#canonical-df84def868cc73d6ffffc568e2956d1cd463d618f8305e6d1e14656bd17948da)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-23cada312bda320b7efc68a617dd2908f73a2b8261f34adfc4988ae22d9199ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0b02239927d4a9e68d8a72fea45347e6c26a9167f17de21f42e4237a71fb253"></a>

## admin_user_credentials — admin_user_credentials / ca2c0b4120e5 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- admin_user_credentials

<a id="canonical-e9ad41bd99ba0698328622f3272de87770af5f50de03f53e1b0b77569e75c599"></a>

Type: `"single"`. Computed.

Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin'
user will be setup and customers can access these nodes via either the node local WebUI or via SSH
to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access.

<a id="canonical-5274a1fed842ed7de2894964fa56cd4209a972cd3e5763f7cdf62123ca03f731"></a>

## Direct properties — admin_user_credentials / ca2c0b4120e5 / 3

- [admin_password](data-sources--site--reference--group-001.md#canonical-14281ce1bce71d7253fe35157074f697cabb0b31626b218742d865e77d59a82d): complete subsection reference.

<a id="canonical-192d670828ae5d42869a91efb1eb542cb85817330626938523cd1bb6b868b96c"></a>

<a id="canonical-149a9e0b0f94f4973806903c6925e13f790caae5f7aa85e717faaef83c5d5cb6"></a>

## ssh_key property — admin_user_credentials / ca2c0b4120e5 / 4

Type: `"string"`. Computed.

Provided Public SSH key can be used for accessing nodes of the site. When provided, customers can
SSH to the nodes of this Customer Edge site using admin as the user.

<a id="canonical-4240323ad1c49e41e63e176e0c222ff32117b9796c13e79cf7a94391ba6df93e"></a>

## Next pages — admin_user_credentials / ca2c0b4120e5 / 5

- [admin_user_credentials.admin_password](data-sources--site--reference--group-001.md#canonical-14281ce1bce71d7253fe35157074f697cabb0b31626b218742d865e77d59a82d)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-14281ce1bce71d7253fe35157074f697cabb0b31626b218742d865e77d59a82d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e10d7d13f31a72b858c870093ecc07d864248496a0796d0b0513a0c802ba82e"></a>

## admin_user_credentials.admin_password — admin_user_credentials.admin_password / f25d10df9a81 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [admin_user_credentials](data-sources--site--reference--group-001.md#canonical-23cada312bda320b7efc68a617dd2908f73a2b8261f34adfc4988ae22d9199ce)
- admin_user_credentials.admin_password

<a id="canonical-b8b3bef1e8d1e9cbecc1d7f866dba33a7fa6389210c54aae7a64ba561c040a53"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

<a id="canonical-47f4071df58c373f14b1438483f1ed128d5fd32bc12c58033c92bee5c2c594b3"></a>

## Direct properties — admin_user_credentials.admin_password / f25d10df9a81 / 3

- [blindfold_secret_info](data-sources--site--reference--group-001.md#canonical-d024cb02519b8127d763f5025f9722d5c1dac3d98a60299b1b9a3db3f8fd820b): complete subsection reference.

- [clear_secret_info](data-sources--site--reference--group-001.md#canonical-62e3aee84b3f8dc5c20f8e6cba1ac30305609d5e762311f0820888f4597e11f4): complete subsection reference.

<a id="canonical-aa4c93b0bce7e59689d7ca18083587f04ca03c69ddf3a9b5996e920b071aebe0"></a>

## Next pages — admin_user_credentials.admin_password / f25d10df9a81 / 4

- [admin_user_credentials.admin_password.blindfold_secret_info](data-sources--site--reference--group-001.md#canonical-d024cb02519b8127d763f5025f9722d5c1dac3d98a60299b1b9a3db3f8fd820b)
- [admin_user_credentials.admin_password.clear_secret_info](data-sources--site--reference--group-001.md#canonical-62e3aee84b3f8dc5c20f8e6cba1ac30305609d5e762311f0820888f4597e11f4)
- [admin_user_credentials](data-sources--site--reference--group-001.md#canonical-23cada312bda320b7efc68a617dd2908f73a2b8261f34adfc4988ae22d9199ce)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-d024cb02519b8127d763f5025f9722d5c1dac3d98a60299b1b9a3db3f8fd820b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-209c93ddac65293e9695327b662541168d3af1b68843d93af84749ad115d4b61"></a>

## admin_user_credentials.admin_password.blindfold_secret_info — admin_user_credentials.admin_password.blindfold_secret_info / bb7172d98eb6 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [admin_user_credentials](data-sources--site--reference--group-001.md#canonical-23cada312bda320b7efc68a617dd2908f73a2b8261f34adfc4988ae22d9199ce)
- [admin_user_credentials.admin_password](data-sources--site--reference--group-001.md#canonical-14281ce1bce71d7253fe35157074f697cabb0b31626b218742d865e77d59a82d)
- admin_user_credentials.admin_password.blindfold_secret_info

<a id="canonical-88ceed83ab3ef478c3f4830fae173388c2e9598eff9866cb8759c495673b43e0"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

<a id="canonical-f3c1eb3c3e2588a58b5a508ad7ec6cd0a05b2a905722d9b594a964986efd58be"></a>

## Direct properties — admin_user_credentials.admin_password.blindfold_secret_info / bb7172d98eb6 / 3

<a id="canonical-9fec5de6cc3a88d55daaccb0a910d15e62440660dd449eeb1f8bcd33b33de212"></a>

<a id="canonical-8733c257d3a48cf8826550ad38d80ba7bcecd1776811eb189ce1cd2c4494f421"></a>

## decryption_provider property — admin_user_credentials.admin_password.blindfold_secret_info / bb7172d98eb6 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

<a id="canonical-90b1a0bd73f0f823a2753e959ccfbc3357050b2844372093dcda19d97a3fec1f"></a>

<a id="canonical-768801b0ff9e8d3a9a74cecd1d829718c1d309c363002ec6bcd7193c82c2cfd4"></a>

## location property — admin_user_credentials.admin_password.blindfold_secret_info / bb7172d98eb6 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

<a id="canonical-3dbc5da04df28793b8be6bb5d80b15157be06159c8bcab4172dfae817a4992e4"></a>

<a id="canonical-64162a7021366b8820f25df609dcd4295c27fa0853e30cd5dcc053b87a5c8b6b"></a>

## store_provider property — admin_user_credentials.admin_password.blindfold_secret_info / bb7172d98eb6 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1f3763d58b2b062b697bdee68b71fee7d8ada0a3118d2ed2bc20fcd3ce951f64"></a>

## Next pages — admin_user_credentials.admin_password.blindfold_secret_info / bb7172d98eb6 / 7

- [admin_user_credentials.admin_password](data-sources--site--reference--group-001.md#canonical-14281ce1bce71d7253fe35157074f697cabb0b31626b218742d865e77d59a82d)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-62e3aee84b3f8dc5c20f8e6cba1ac30305609d5e762311f0820888f4597e11f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a93b4b780b8eb678133f3dc4882761b97f2626cfb31187a79f921aded7a67656"></a>

## admin_user_credentials.admin_password.clear_secret_info — admin_user_credentials.admin_password.clear_secret_info / 483c48944dae / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [admin_user_credentials](data-sources--site--reference--group-001.md#canonical-23cada312bda320b7efc68a617dd2908f73a2b8261f34adfc4988ae22d9199ce)
- [admin_user_credentials.admin_password](data-sources--site--reference--group-001.md#canonical-14281ce1bce71d7253fe35157074f697cabb0b31626b218742d865e77d59a82d)
- admin_user_credentials.admin_password.clear_secret_info

<a id="canonical-c14b3d6573633b52179abc4ae97c8eec1f6fabc9009487e4365ca473d9dafe00"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

<a id="canonical-34cfef6e3272a0a299be050392a0d9184eb4afe9c3037885849add22db519664"></a>

## Direct properties — admin_user_credentials.admin_password.clear_secret_info / 483c48944dae / 3

<a id="canonical-2033164fd940b85d87f97e7925f666ff6856a8d724c0fad265c5ad673ffd98db"></a>

<a id="canonical-8d6e879dd5fcd78529c9453d8acb299b409f466e84c36b6e5f669adc3ec46e71"></a>

## provider_ref property — admin_user_credentials.admin_password.clear_secret_info / 483c48944dae / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-36daa3f0efc1ce8c564a9dfea21e13851ff10ad42515a779206837f0bf860028"></a>

<a id="canonical-16bf9de96cfcfd0d8ef65b6eb8095da2cd375aa98aba75762ebcf602c1989d28"></a>

## url property — admin_user_credentials.admin_password.clear_secret_info / 483c48944dae / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

<a id="canonical-95e8b2ea886521f4187e03579dfbafd9bf8aaa3d35171100da66b0bd5d2eddf9"></a>

## Next pages — admin_user_credentials.admin_password.clear_secret_info / 483c48944dae / 6

- [admin_user_credentials.admin_password](data-sources--site--reference--group-001.md#canonical-14281ce1bce71d7253fe35157074f697cabb0b31626b218742d865e77d59a82d)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-85706de1aed9b2978a14e030cedde0be433104392d6a8841ea76bb2dfbeaa589"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-001d44dd068192e3075364fe423b769e7fc34a8bb903b0eae4b31c2c56950d8d"></a>

## connected_re — connected_re / 8c29edf84974 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- connected_re

<a id="canonical-a7d91e199cdbd33164b316d6de5f60318841d45d797729c40eb6e444800ae63d"></a>

Type: `"list"`. Computed.

Following fields are only for customer edge sites List of REs to which to which this CE initiates
IPsec/SSL connection to.

<a id="canonical-586336a80f10c3394781df31a0c16242437585220fb06c756ce7a6595706e12a"></a>

## Direct properties — connected_re / 8c29edf84974 / 3

<a id="canonical-4e21f5cc085189e431aa797b96ed0d73a4fe8207ed18def6de422d1be826c3b6"></a>

<a id="canonical-6ab17a8da48d6ff05218a7902f6f5d8cd3915aa3bf57c2e01ec504f13e83c7a7"></a>

## kind property — connected_re / 8c29edf84974 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="canonical-aff8fca639bddead506db574f7a6ce6220b4732b11cd94457e8e32fd28dc84fc"></a>

<a id="canonical-39791b2682d08876cb0ad75fd909914e180700d3489d5b63f09c016057eaf55e"></a>

## name property — connected_re / 8c29edf84974 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="canonical-fe615b7fd5875de442912cb7e5a07567bea16475275136dd33bc5ada5520d2e9"></a>

<a id="canonical-c6eac0e908732b9b072724d31e2f0dc7c3abc269cdb7385d17b01fa0001e782e"></a>

## namespace property — connected_re / 8c29edf84974 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="canonical-a8c28fd0d47b7670bcc7a91a2a4ee11afc565887537c04a67d70c210f63a10ca"></a>

<a id="canonical-54ea0ae8e751e32d6d1607a3fb7f9f2653671f4a1dbaf5a23b476b5b8aa7fafc"></a>

## tenant property — connected_re / 8c29edf84974 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="canonical-7b756f810be1acdadc0b89d87eb14c1016477bd2593b1b84387d94d5a9e941c3"></a>

<a id="canonical-ca0abbd267799327eebb2d841cd8bd26c4775727a4dc313345366c810bd67b9a"></a>

## uid property — connected_re / 8c29edf84974 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

<a id="canonical-b2b2b9e7ca95979ca046aaf78e9dd112759b10e8d7932e5418593a8f479d5b31"></a>

## Next pages — connected_re / 8c29edf84974 / 9

- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-70617853cc03e32ab5d93d5623a6b5794c6f8b550e5f402d8a1445c8bba394a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cb7755769172aa6f13602c9651bdfe44aa7580da5a3ae0a04064789c6744e26"></a>

## connected_re_for_config — connected_re_for_config / 2ee3e16e746c / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- connected_re_for_config

<a id="canonical-97421362f29939bb8ef71a2031f1be9bbb20c90f0f51cfa1101f50771490e494"></a>

Type: `"list"`. Computed.

Valid only for CE site object List of REs which can send config to this CE site.

<a id="canonical-e44461b8a284d5f223f30e18fd86ac8a953291d44b854ca390f0055e2c14ee76"></a>

## Direct properties — connected_re_for_config / 2ee3e16e746c / 3

<a id="canonical-f273285434b4b64d3c26f09179625ee704bb055e02bf77387de6afd347aef11d"></a>

<a id="canonical-77551cc606b17bb1e970b33dfdede4349c50cdaf657aada0657448eaec7ade3c"></a>

## kind property — connected_re_for_config / 2ee3e16e746c / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="canonical-76e16a31ea9bdc0aae6ea34f9772ee6230bf9cd6d22a95d4ac2e6479eb2dfaee"></a>

<a id="canonical-468b41bd96f62de716fc27ba400d71d5b96ad60228fc2e265778a589ec9231a9"></a>

## name property — connected_re_for_config / 2ee3e16e746c / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="canonical-9c7ddb44bd40ffa4b50997bc591082f41b16152478f3a82a85672044d3b3fa87"></a>

<a id="canonical-4c2ac0837c29e66ac55fff9510b7f9ad8e213cb6febdb8925fcb2d2c577bc131"></a>

## namespace property — connected_re_for_config / 2ee3e16e746c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

<a id="canonical-4daa101ef9a3786128e1b82b3dc91b94442d1420ae9dcda7a653261fb6a22eb4"></a>

<a id="canonical-7819a377f63875c93a8ce3bf9aea3dc9d6e099afc10a1077e06d618d35f6a037"></a>

## tenant property — connected_re_for_config / 2ee3e16e746c / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="canonical-41f6e3f5f5f1a2c2e279508d3ada0b7db91bfa33c50ac1118e404fbd4c0205d4"></a>

<a id="canonical-eefeb7482745ad2d3919db9d8dd3bad2084b891a92c27b9e7564e305b9b78d60"></a>

## uid property — connected_re_for_config / 2ee3e16e746c / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

<a id="canonical-7dad534a10b8dee15508cff9e4b3e44703e32a34929b2c06bb48333d91c94a3b"></a>

## Next pages — connected_re_for_config / 2ee3e16e746c / 9

- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-4acdcf69b09e0ea8e1684d40f9537429ddefa1d373e5328e2ea7fc7e4f2f9faf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d4675188028a3a07ef0cee07d83202d15e09949d54a89155b501b74f5c8869f"></a>

## coordinates — coordinates / 259e297b2744 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- coordinates

<a id="canonical-03b01fed26df694dbf42e985da27d2e75b3f306656a5bdd4816ea426cf76c736"></a>

Type: `"single"`. Computed.

Coordinates of the site which provides the site physical location.

<a id="canonical-c633d1b2bf3c2b1300a9b91faff022729896ea2bc9fde6035f167522170fb65f"></a>

## Direct properties — coordinates / 259e297b2744 / 3

<a id="canonical-0bfe4b6ae704a43dd1ae9dc136d9477ea48e2a846916d4861bfbf9cacebde14c"></a>

<a id="canonical-6bd143467efb68a469640a121b4a7f8364466177eeb9a83d9b178e2720de62cc"></a>

## latitude property — coordinates / 259e297b2744 / 4

Type: `"number"`. Computed.

Latitude. Latitude of the site location.

<a id="canonical-22a44e5606705a32814ee60eb588b4a5884942c98b850bcaea32d9b126a757ae"></a>

<a id="canonical-a5cb23d9ee71cf0249425aa394eea0cf1cf0b2b89b448981d3c6ac27aa4fd850"></a>

## longitude property — coordinates / 259e297b2744 / 5

Type: `"number"`. Computed.

Longitude. Longitude of site location.

<a id="canonical-05a6fffe1eb7fe136efd8582a486625ad607dc10c564b8e82dd56eed624fd3fa"></a>

## Next pages — coordinates / 259e297b2744 / 6

- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-bd17269542aad6bf8548924856fd769c4a855887d702aba855aa7688e63cf3f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0ae8d79b5f6ef382a2e1a98775b9dc0a9bd246809200c8ee4d8a14c06a419a9"></a>

## default_underlay_network — default_underlay_network / 248947760135 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- default_underlay_network

<a id="canonical-a762a1a85bca5dbefaab87be71767cb764b9abada86afc9b6176b443cea0e0e8"></a>

Type: `"single"`. Computed.

Optional, virtual network to be used as underlay for different overlay protocols (SRv6, IP-in-IP
tunnels for DC Cluster Group) Default is site-local-outside network.

<a id="canonical-41818ff3358b7f916a9bb79b5dd8192992ace84ec02df38c347ce478b2e8f3c0"></a>

## Direct properties — default_underlay_network / 248947760135 / 3

- [site_local_inside](data-sources--site--reference--group-001.md#canonical-c53b690d07a2f0544045917d2d636e25819592d562dbee4db3193e96ce0bdc16): complete subsection reference.

- [site_local_outside](data-sources--site--reference--group-001.md#canonical-edf2dc101fc413ccf9719469478f853a6b9ee6be95908cd501c9d08f02df72c2): complete subsection reference.

<a id="canonical-f6c33b339dacb389149ff3918bfe1f6cfa93889830aa6dcb17bcea9ad3a28d4f"></a>

## Next pages — default_underlay_network / 248947760135 / 4

- [default_underlay_network.site_local_inside](data-sources--site--reference--group-001.md#canonical-c53b690d07a2f0544045917d2d636e25819592d562dbee4db3193e96ce0bdc16)
- [default_underlay_network.site_local_outside](data-sources--site--reference--group-001.md#canonical-edf2dc101fc413ccf9719469478f853a6b9ee6be95908cd501c9d08f02df72c2)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-c53b690d07a2f0544045917d2d636e25819592d562dbee4db3193e96ce0bdc16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fc94c708df42b7ba704c7baa9c127842ac3b78b6ad50e2bf6d49497338afefe"></a>

## default_underlay_network.site_local_inside — default_underlay_network.site_local_inside / 0f83f95cad18 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [default_underlay_network](data-sources--site--reference--group-001.md#canonical-bd17269542aad6bf8548924856fd769c4a855887d702aba855aa7688e63cf3f5)
- default_underlay_network.site_local_inside

<a id="canonical-fe41799b0711b493e75b835e956c44a1b6395b60d9eb2ea09fd18d23c23065dc"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-a39b977c924369289c3f1dbc1c638676a8fb5dca2b10d90ab38cb7a7ee4325df"></a>

## Direct properties — default_underlay_network.site_local_inside / 0f83f95cad18 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47c998d9440fdf456d03f6100756dd99c8962558326321ba54c1cc6d215f42eb"></a>

## Next pages — default_underlay_network.site_local_inside / 0f83f95cad18 / 4

- [default_underlay_network](data-sources--site--reference--group-001.md#canonical-bd17269542aad6bf8548924856fd769c4a855887d702aba855aa7688e63cf3f5)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-edf2dc101fc413ccf9719469478f853a6b9ee6be95908cd501c9d08f02df72c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-011a961210d0788ae6221bbbcc17b497e78fc66037588493a10eb325fb095def"></a>

## default_underlay_network.site_local_outside — default_underlay_network.site_local_outside / 726486fbdfc6 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [default_underlay_network](data-sources--site--reference--group-001.md#canonical-bd17269542aad6bf8548924856fd769c4a855887d702aba855aa7688e63cf3f5)
- default_underlay_network.site_local_outside

<a id="canonical-1253b80574300d4d0a3854f83dcd57f7f85d82d342be942f2c2cb258fb98e627"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-fbc5f1ab08b5e5efcbc0fa019b96f628921eb4ea5e444a2cd6240a73fe211eac"></a>

## Direct properties — default_underlay_network.site_local_outside / 726486fbdfc6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9f0d1fe11f71547e89504df1865c467198bd2a9e53ee4dd2f95bcabbb7727dd"></a>

## Next pages — default_underlay_network.site_local_outside / 726486fbdfc6 / 4

- [default_underlay_network](data-sources--site--reference--group-001.md#canonical-bd17269542aad6bf8548924856fd769c4a855887d702aba855aa7688e63cf3f5)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-82ce9fbbef5339fa3591881ed6c4c9c12363dc591088c814863b0e183f8ad380"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b842e292d01850b6f96da37b3b4596e29f0f777e02a59d4e87c4b3557e7332b0"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 595e4ab1ef24 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- kubernetes_upgrade_drain

<a id="canonical-71d124a4030aacd174f374d619778ec41f397171cd6a34b74db40be4f074075c"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

<a id="canonical-264ccfad98aabdcf910472926a976f89805c2d4dc9b45f7c3fd3430e1d860e57"></a>

## Direct properties — kubernetes_upgrade_drain / 595e4ab1ef24 / 3

- [disable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-ca13617df8a83677d446be7c5da5ae092ccbc859c9108636bb38e5cb8962d002): complete subsection reference.

- [enable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-430d1b9ac119b46e47b9b5d5eca859876eb9bf01baf6a028c5384dc1997894b7): complete subsection reference.

<a id="canonical-1d07f5d8c3c222fec85178bab6248433140becc90bfd381d1ded4dad510e2d62"></a>

## Next pages — kubernetes_upgrade_drain / 595e4ab1ef24 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-ca13617df8a83677d446be7c5da5ae092ccbc859c9108636bb38e5cb8962d002)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-430d1b9ac119b46e47b9b5d5eca859876eb9bf01baf6a028c5384dc1997894b7)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-ca13617df8a83677d446be7c5da5ae092ccbc859c9108636bb38e5cb8962d002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cbe7f175d6ec27f91d15a111ffe4ec339dbfa3ebd5818d9032f39d7969bfa73"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — kubernetes_upgrade_drain.disable_upgrade_drain / 925f6b2062bf / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [kubernetes_upgrade_drain](data-sources--site--reference--group-001.md#canonical-82ce9fbbef5339fa3591881ed6c4c9c12363dc591088c814863b0e183f8ad380)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-2d608f33760dea286591a976a58d6a23adf9fd3459f458ae8b007c5cdebbde85"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable upgrade drain.

<a id="canonical-90adb9fc1e0a2b80093a9aa3b6d80d1f3f4dd1a1b00744a7ba1e2abe93431133"></a>

## Direct properties — kubernetes_upgrade_drain.disable_upgrade_drain / 925f6b2062bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe1a8890b4d91a9d1daa0928455d18a07f4789acc8313fe8cf2c75d06e0ff4d2"></a>

## Next pages — kubernetes_upgrade_drain.disable_upgrade_drain / 925f6b2062bf / 4

- [kubernetes_upgrade_drain](data-sources--site--reference--group-001.md#canonical-82ce9fbbef5339fa3591881ed6c4c9c12363dc591088c814863b0e183f8ad380)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-430d1b9ac119b46e47b9b5d5eca859876eb9bf01baf6a028c5384dc1997894b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75911ef22c23e8c1bd065c85bb0fddaec5ef926965a708801ab72927e0ceb75a"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — kubernetes_upgrade_drain.enable_upgrade_drain / 17364751fdf1 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [kubernetes_upgrade_drain](data-sources--site--reference--group-001.md#canonical-82ce9fbbef5339fa3591881ed6c4c9c12363dc591088c814863b0e183f8ad380)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-3a5f0d0538e4764a35eb84d26a4a96346bd8c93beea61011f9e1746868cc114e"></a>

Type: `"single"`. Computed.

Specify batch upgrade settings for worker nodes within a site.

<a id="canonical-f0e51593439e44c06065a60ea1d224ab6f30eae1bebd01b294c24cd9f0dc4fba"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain / 17364751fdf1 / 3

- [disable_vega_upgrade_mode](data-sources--site--reference--group-001.md#canonical-7ef531d25473cd72819eeb0538adbfe6a417f7559226c5c97bee8689b757a761): complete subsection reference.

<a id="canonical-bdb3194819542b209c7db0bcad89876c8a83d49de55a959a3d6da57701b9e058"></a>

<a id="canonical-fe9a280a08a43c7655a7c008feadfb7c6976c02627dae9c4d6c9ba2d9c41a2cd"></a>

## drain_max_unavailable_node_count property — kubernetes_upgrade_drain.enable_upgrade_drain / 17364751fdf1 / 4

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

<a id="canonical-84f262c0a2973f3e5431e5c78d5b51725eb2aada3e8427dc26bcfec9b934481b"></a>

<a id="canonical-93ad187edf032404242831ae16ebea267190e3630aa4137113fd8d55ed1f63f7"></a>

## drain_max_unavailable_node_percentage property — kubernetes_upgrade_drain.enable_upgrade_drain / 17364751fdf1 / 5

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-bdfa3c7c4b9c935408578410a4fc3b402eefb161fc46c49f29440218b6b43aa7"></a>

<a id="canonical-718e705f34b1418e83516b734c53a3285c0b95dbed6d11c1386b41a497e2765c"></a>

## drain_node_timeout property — kubernetes_upgrade_drain.enable_upgrade_drain / 17364751fdf1 / 6

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

- [enable_vega_upgrade_mode](data-sources--site--reference--group-001.md#canonical-7bd7b26ed43a006e791c0da1f5f27aee208e676808311b8472454d6e3a8bb441): complete subsection reference.

<a id="canonical-e9aa6842406dc79908feb87f43fe6f99a9b0e84842f36d3157d64b5291d25872"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain / 17364751fdf1 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--site--reference--group-001.md#canonical-7ef531d25473cd72819eeb0538adbfe6a417f7559226c5c97bee8689b757a761)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--site--reference--group-001.md#canonical-7bd7b26ed43a006e791c0da1f5f27aee208e676808311b8472454d6e3a8bb441)
- [kubernetes_upgrade_drain](data-sources--site--reference--group-001.md#canonical-82ce9fbbef5339fa3591881ed6c4c9c12363dc591088c814863b0e183f8ad380)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-7ef531d25473cd72819eeb0538adbfe6a417f7559226c5c97bee8689b757a761"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b320e7f6b1be8fbc66d28a7a24362954a945b658b24c193f3d8df022cbb72b66"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 8519861bf1d7 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [kubernetes_upgrade_drain](data-sources--site--reference--group-001.md#canonical-82ce9fbbef5339fa3591881ed6c4c9c12363dc591088c814863b0e183f8ad380)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-430d1b9ac119b46e47b9b5d5eca859876eb9bf01baf6a028c5384dc1997894b7)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-d02e64fe3052930a946b69e9812d31b4c5a6ae36720436ca682e4a30fc18bb9e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable vega upgrade mode.

<a id="canonical-49c4d0c5fe8703d60f41635a01eef4c7403d4a7ce6fc25ece18a46c3bed8ea52"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 8519861bf1d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b855123b7d05a409997280b651e61dd593825a7433364435978bcee2dd24c22a"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 8519861bf1d7 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-430d1b9ac119b46e47b9b5d5eca859876eb9bf01baf6a028c5384dc1997894b7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-7bd7b26ed43a006e791c0da1f5f27aee208e676808311b8472454d6e3a8bb441"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38d9d522bda59d6493e72c6727c3c8f0d99220ac56a0a9f6957866c7d0503013"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 5ce59efcfda6 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [kubernetes_upgrade_drain](data-sources--site--reference--group-001.md#canonical-82ce9fbbef5339fa3591881ed6c4c9c12363dc591088c814863b0e183f8ad380)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-430d1b9ac119b46e47b9b5d5eca859876eb9bf01baf6a028c5384dc1997894b7)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-7f4cd3224ac315d2989cfc1ac3891656dee07d004b5fc2c4997ec9c32cb143f0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable vega upgrade mode.

<a id="canonical-96813e710bc7a6207032084c9ad092e93b60972f53cd3bf6d07f44fac2509b8c"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 5ce59efcfda6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1018d271a965b5f084ea2ce3e6c004e417aee80df18716ff3011ce5fc7bcd0a4"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 5ce59efcfda6 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--site--reference--group-001.md#canonical-430d1b9ac119b46e47b9b5d5eca859876eb9bf01baf6a028c5384dc1997894b7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-e659811d82fb9727fa18dff418f5ce4c62d4b56265159514aabefeaf9e7885da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3de74901a26a5008092f64681aa7781f01c2439224ae81666349cb0990ba40c0"></a>

## main_nodes — main_nodes / b00734719a44 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- main_nodes

<a id="canonical-756f1e1a026215d585567ac0c026d9eed0d8ad0bc9f5c0f73845313eddbc3d14"></a>

Type: `"list"`. Computed.

Connectivity information of main/master nodes to create a full mesh of Phobos services across all
CEs in a site-mesh-group or dc-cluster-group.

<a id="canonical-3ffb4a6e4beb3d4fb3fb1acebf8fdbd56cefe1ef64dd7bfe519553aafd8da9d1"></a>

## Direct properties — main_nodes / b00734719a44 / 3

<a id="canonical-e3552c5824655418a6ab2bb64e153227f3a6e9857a63370e10f4d0a96f62d5d7"></a>

<a id="canonical-e450fac3e2ff8ae9841858ea12549c0627eeb5f1fdacf895d8fd9336a0521614"></a>

## name property — main_nodes / b00734719a44 / 4

Type: `"string"`. Computed.

Name of the master/main node on the site.

<a id="canonical-c28e28e21851381098e47ff4a1e201b089dc920b7b14c1246bb62f9e3a1f5967"></a>

<a id="canonical-14436289c4e5b6cde170f0bd29b38454f561105173dea7b3aa2a00d05995264b"></a>

## sli_address property — main_nodes / b00734719a44 / 5

Type: `"string"`. Computed.

Site Local Inside IP addresses. Site Local Inside IP address.

<a id="canonical-47bb30fed57cc0f141d5fa281baa4842037b2f074cb937a6cc0b73d226074729"></a>

<a id="canonical-9f651c65f5d37dab9ddd68b842d4820a9a922fc970c879ba60402d72274d304f"></a>

## slo_address property — main_nodes / b00734719a44 / 6

Type: `"string"`. Computed.

Site Local Outside IP addresses. Site Local Outside IP address.

<a id="canonical-164fe5508f18c8c0a87aa4e124db587ec58d0b25f587e1a3a22ac405e353301a"></a>

## Next pages — main_nodes / b00734719a44 / 7

- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-c156e10b94775d615e04b930cd431c283b2f53cf610d0eb11028b58aa65a8c55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45d82b37f402bb509eca3d35a19995c20b15617a3a215d69e6e1ac0c6c8c6350"></a>

## private_connectivity — private_connectivity / 5f3f2a894878 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- private_connectivity

<a id="canonical-443c7f09d056cf32a0c1f563d105c01a3bbb2a9676b07da28e5c6c0b4252db63"></a>

Type: `"single"`. Computed.

Private Connectivity Information like ADN network name and cloud link information.

<a id="canonical-161606edddaf54fa71353a7d4c9aec4bbbfa4231f5bf5865c4196366e50f28e7"></a>

## Direct properties — private_connectivity / 5f3f2a894878 / 3

- [cloud_link](data-sources--site--reference--group-001.md#canonical-fa4ff091cf94068909ed95a1cf7f77d2565b2a4d745ed6b91782ee04d701bc18): complete subsection reference.

<a id="canonical-6bb2f74a9d0f645c0e6987a1352c46becc1351789a9e699524e624dd642f7098"></a>

<a id="canonical-29a6558eea567136d081b2d545589711fee9fa28bde7101d68460863135b3667"></a>

## private_network_name property — private_connectivity / 5f3f2a894878 / 4

Type: `"string"`. Computed.

ADN Network Name for private access connectivity to F5XC ADN.

<a id="canonical-f68dcddb65e4607f0c0aa323f37a1bc79c78812a2a3903d53a99379fa1f5de77"></a>

## Next pages — private_connectivity / 5f3f2a894878 / 5

- [private_connectivity.cloud_link](data-sources--site--reference--group-001.md#canonical-fa4ff091cf94068909ed95a1cf7f77d2565b2a4d745ed6b91782ee04d701bc18)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-fa4ff091cf94068909ed95a1cf7f77d2565b2a4d745ed6b91782ee04d701bc18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35a34540500ab9acd624680b8d69aa2b7b875dfd3e61b305b1b1d718e9425840"></a>

## private_connectivity.cloud_link — private_connectivity.cloud_link / 6cc4e0789b83 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [private_connectivity](data-sources--site--reference--group-001.md#canonical-c156e10b94775d615e04b930cd431c283b2f53cf610d0eb11028b58aa65a8c55)
- private_connectivity.cloud_link

<a id="canonical-884826ddd9549ed200f72212ce746e3d098b33b597fe7a6f44e3d14149c81ca3"></a>

Type: `"single"`. Computed.

Information related to cloud link used by the site.

<a id="canonical-917c04c957cff44abf04ae9d249c47431d8e36f759ce76e4e71aae6f5bfa01ca"></a>

## Direct properties — private_connectivity.cloud_link / 6cc4e0789b83 / 3

<a id="canonical-16385b40c01df2dd292b41761af77681a9599e8d4e9310c04c85ea14ade42fe5"></a>

<a id="canonical-f195e8ea78b831fca3de6560b29b3d15eada9acd5f4baa6784abc6df5212c2e1"></a>

## name property — private_connectivity.cloud_link / 6cc4e0789b83 / 4

Type: `"string"`. Computed.

Name of the the CloudLink used with this site.

<a id="canonical-18c57637c245572d111dfc52d7c50378a084ff39f4bd521b3354c0e360832a97"></a>

<a id="canonical-2ff1da39c77b3113c01758a85e5cb747893427edc68753678ea244aa45087293"></a>

## state property — private_connectivity.cloud_link / 6cc4e0789b83 / 5

Type: `"string"`. Computed.

\[Enum: UP|DOWN|DEGRADED|NOT\_APPLICABLE\] State of the CloudLink connections - UP: Up CloudLink and
their corresponding Direct Connect connections are up and healthy - DOWN: Down CloudLink and their
corresponding Direct Connect connections are down - DEGRADED: Degraded Some of Direct Connect
connections with the CloudLink are down .. Possible values are \`UP\`, \`DOWN\`, \`DEGRADED\`,
\`NOT\_APPLICABLE\`. Defaults to \`UP\`.

<a id="canonical-616527c61aa89e78b6f61db7ec97c23d712605919b4957c89503f56c4ee661d1"></a>

## Next pages — private_connectivity.cloud_link / 6cc4e0789b83 / 6

- [private_connectivity](data-sources--site--reference--group-001.md#canonical-c156e10b94775d615e04b930cd431c283b2f53cf610d0eb11028b58aa65a8c55)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-b1e08c929b5c34e6b18d034595249b5dd6e81bb071980608b29fbe99c33a6fcf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5c0bcdb23f05a2b3a6da9cd5bd6b78c5718427f164b6f3790180e3f3ec40af7"></a>

## re_select — re_select / 4107c10d9067 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- re_select

<a id="canonical-70af337bc649fc0197600ab7da9df9b1b426de15d2afa4c8212573ca0f8d4d88"></a>

Type: `"single"`. Computed.

Selection criteria to connect the site with F5 Distributed Cloud Regional Edge(s).

<a id="canonical-a0d46a0a3ca2b1e73b7ceb2aa894cab677d8a80546fa1f7b92968341611958c3"></a>

## Direct properties — re_select / 4107c10d9067 / 3

- [geo_proximity](data-sources--site--reference--group-001.md#canonical-10697e7fd4b237fcef69b0318bae85b7915929089c33981884ae7f73ef6d19f2): complete subsection reference.

<a id="canonical-207ca62280fb3c03ef9d3ac13e323a8b56fa4927562805beaa06ec7b7a09a9d7"></a>

<a id="canonical-74b85430b108139de7ff53072072663485d163869b2b7bdb85fdfe43b0d22b01"></a>

## specific_geography property — re_select / 4107c10d9067 / 4

Type: `"string"`. Computed.

Geographic selection for the site's Regional Edge connections.

- [specific_re](data-sources--site--reference--group-001.md#canonical-bc13544c9925f4dc422bf5332c386ecfdd85f08790ad987547f455ff5ee54cfc): complete subsection reference.

<a id="canonical-ab00c55acc50067d040ffac7e559f6cbe3952a04c7466c4d1538fb3bc0d4bd58"></a>

## Next pages — re_select / 4107c10d9067 / 5

- [re_select.geo_proximity](data-sources--site--reference--group-001.md#canonical-10697e7fd4b237fcef69b0318bae85b7915929089c33981884ae7f73ef6d19f2)
- [re_select.specific_re](data-sources--site--reference--group-001.md#canonical-bc13544c9925f4dc422bf5332c386ecfdd85f08790ad987547f455ff5ee54cfc)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-10697e7fd4b237fcef69b0318bae85b7915929089c33981884ae7f73ef6d19f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb0e70d132ad7abd9930a16423cc65840af159c3ff1bb3bf24962baf2cd4b530"></a>

## re_select.geo_proximity — re_select.geo_proximity / 5d9054224b35 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [re_select](data-sources--site--reference--group-001.md#canonical-b1e08c929b5c34e6b18d034595249b5dd6e81bb071980608b29fbe99c33a6fcf)
- re_select.geo_proximity

<a id="canonical-66c8c4bb05e8f25cf773729131aa627c4eddb5ff59cb6be18590475af11804c0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for geo proximity.

<a id="canonical-7e2fefe76ea807b7de5728c01b5b25c3d8f923736fad61ac10544ac50d7e147f"></a>

## Direct properties — re_select.geo_proximity / 5d9054224b35 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-710b1b8821908644ce49dfa7b2136d374a59c11e3f09fc6332421813f1a3cdf0"></a>

## Next pages — re_select.geo_proximity / 5d9054224b35 / 4

- [re_select](data-sources--site--reference--group-001.md#canonical-b1e08c929b5c34e6b18d034595249b5dd6e81bb071980608b29fbe99c33a6fcf)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-bc13544c9925f4dc422bf5332c386ecfdd85f08790ad987547f455ff5ee54cfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4db68423c9bb56b6421b1acab86cdd9272c976d20132871f016c8e9c9e15401"></a>

## re_select.specific_re — re_select.specific_re / 34b331f45a46 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [re_select](data-sources--site--reference--group-001.md#canonical-b1e08c929b5c34e6b18d034595249b5dd6e81bb071980608b29fbe99c33a6fcf)
- re_select.specific_re

<a id="canonical-7e8edcd4b37f7b746b9b801b761e6fb6e97f29e9797ad49c3b67f84b94b24a7b"></a>

Type: `"single"`. Computed.

Select specific REs. This is useful when a site needs to deterministically connect to a set of REs.
A site will always be connected to 2 REs.

<a id="canonical-35cc57e7ea65f5eca41d26233df5b66b84a7b1d2c9bdcf32be74ff92663e285e"></a>

## Direct properties — re_select.specific_re / 34b331f45a46 / 3

<a id="canonical-3ccf72c32648e873c90685086acd852d65656fb3fdaee62ecd9339a4108fd037"></a>

<a id="canonical-8ece63d4dcefb215f1b7064aff2c5897b44f89d49c3689b0e83db9c05fe6f9d5"></a>

## backup_re property — re_select.specific_re / 34b331f45a46 / 4

Type: `"string"`. Computed.

Select backup RE for this site, cannot be the same as Primary RE.

<a id="canonical-e008d4f98f96d72fb9c58b96aeee8c7dab95d20bac2055ecf8a823bc3b629be5"></a>

<a id="canonical-6963efa47a408a541ac7ea324f7b0b9e623884792e2853c9c0f6f6e2e1f60cfe"></a>

## primary_re property — re_select.specific_re / 34b331f45a46 / 5

Type: `"string"`. Computed.

Primary RE Geography. Select primary RE for this site.

<a id="canonical-3ae14db0bc3a749886be8275ba24529a6b69ee2b9915b875171228e04772ff40"></a>

## Next pages — re_select.specific_re / 34b331f45a46 / 6

- [re_select](data-sources--site--reference--group-001.md#canonical-b1e08c929b5c34e6b18d034595249b5dd6e81bb071980608b29fbe99c33a6fcf)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)

<a id="canonical-df84def868cc73d6ffffc568e2956d1cd463d618f8305e6d1e14656bd17948da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16e95d98cda368c090c0f29ab479fe28268dd54a4624f70328d5e746fde71bf7"></a>

## vip_params_per_az — vip_params_per_az / a91dafe5a3b4 / 2

Breadcrumbs:

- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- vip_params_per_az

<a id="canonical-975f0278fcff1c3020efb711819a9649dbd67a542f1323be97b894196510cc61"></a>

Type: `"list"`. Computed.

Optional Publish VIP Parameters Per AZ for public cloud sites. See documentation for 'VIP' in
advertise policy to see when Inside VIP or Outside VIP is used. When configured, the VIP(s) defined
will be used to publish to external systems like K8s, Consul.

<a id="canonical-c1c46ad40995a793d00d0a190bb3f3508d570c296c10fc821734bb35cb2a3a9b"></a>

## Direct properties — vip_params_per_az / a91dafe5a3b4 / 3

<a id="canonical-1f6a6d15bcd16477a7ead7934dc1921cae6bf26745769cd0ec5470f2cf07caa9"></a>

<a id="canonical-f115aba2fd05acdafabfc57fb7fdded65a590a5ea3a5d7c4c1560973a6e6973e"></a>

## az_name property — vip_params_per_az / a91dafe5a3b4 / 4

Type: `"string"`. Computed.

AZ Name. Name of the Availability zone.

<a id="canonical-116cafdd64a3ad1a3148440312334f88460ff8f1c845e6434aaea6d95ef69413"></a>

<a id="canonical-4efd5df04688ba46d13193f15d213dff1b8e310b8647353103120c4610562e64"></a>

## inside_vip property — vip_params_per_az / a91dafe5a3b4 / 5

Type: `["list", "string"]`. Computed.

Inside VIP(s). List of Inside VIPs for an AZ.

<a id="canonical-48b73b24cd3401657f72e99bf5452ffa21824f42bbd8cf4032e8921dbd8740ba"></a>

<a id="canonical-bce71f78c96a399d4eb5e217b905a82554857c9bac9f76e4426b6a635d47ced8"></a>

## inside_vip_cname property — vip_params_per_az / a91dafe5a3b4 / 6

Type: `"string"`. Computed.

CNAME value for the inside VIP, These are usually public cloud generated CNAME.

<a id="canonical-5d8ec8c07494b3c2bd4cbe454fddd165a2ae83c75a562e792ce81e7d5e2df967"></a>

<a id="canonical-349265228c4af380d0bc0ebd990c197e07dd9706fce5378d1893953e8d300849"></a>

## inside_vip_v6 property — vip_params_per_az / a91dafe5a3b4 / 7

Type: `["list", "string"]`. Computed.

Optional list of Inside IPv6 VIPs for an AZ.

<a id="canonical-890522a3682f28c289c6a77fd10dd9f835634be8486588b6704dae66cfdf45ca"></a>

<a id="canonical-a33de05972e60ad5e794c8e727291b63e9b8faab5a3c44a371858ef22e5d3182"></a>

## outside_vip property — vip_params_per_az / a91dafe5a3b4 / 8

Type: `["list", "string"]`. Computed.

Outside VIP(s). List of Outside VIPs for an AZ.

<a id="canonical-45fb01b0f79b4ae640273c78e22d5c4a98e8657238a5b56547c3e48e2daf6c68"></a>

<a id="canonical-8886eac8bac78e1a8aa298fe68ee7da57c2253c185905dd7aaebfe9e79f4b08c"></a>

## outside_vip_cname property — vip_params_per_az / a91dafe5a3b4 / 9

Type: `"string"`. Computed.

CNAME value for the outside VIP These are usually public cloud generated CNAME.

<a id="canonical-ff7bc32c2422d83ce0bf8a126e4341f51e2bbdb1aeb15b516f015fc91e71014d"></a>

<a id="canonical-5a1280ac6e176b832b7399813a2a2fde03a01a36ba696ba8cea54ceb2e5b8914"></a>

## outside_vip_v6 property — vip_params_per_az / a91dafe5a3b4 / 10

Type: `["list", "string"]`. Computed.

Optional list of Outside IPv6 VIPs for an AZ.

<a id="canonical-88eb2ee63aa464571f759314d866ca3f525bf854a03b54dd45ac4977a7b75efe"></a>

## Next pages — vip_params_per_az / a91dafe5a3b4 / 11

- [Property reference](data-sources--site--reference--group-001.md#canonical-f57c5c2c39acd811d35d80bc061dfe539cb37c8fb15385380fc3d10e36c434f7)
- [xcsh_site](../data-sources/site.md#canonical-d365b726b544f9ad32e25b248398f21a1481685e6b24ca8dd76dca1c8253f1d8)
