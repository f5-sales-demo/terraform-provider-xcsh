---
page_title: "xcsh_certified_hardware reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_certified_hardware reference."
---

# xcsh_certified_hardware reference

<a id="canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07d285d0dbbe41c3543a07d4c2066177349f178fbd19cfd4452c6047d34b31d5"></a>

## Property reference — Property reference / c22b4a524092 / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- Property reference

<a id="canonical-d9bb2b2daa9943fcf443e0c564e10afedf30278733274ac7294163cd80382119"></a>

## Direct properties — Property reference / c22b4a524092 / 3

<a id="canonical-4d08f3e105e21129a0f84c48563515ab3d62c207500e5ac91c4bab01afaf896c"></a>

<a id="canonical-c4c5db98950cf47ff42e1879ec112b1642080b7cec8bac6bb899b33bf82edf01"></a>

## annotations property — Property reference / c22b4a524092 / 4

Type: `["map", "string"]`. Computed.

Annotations.

<a id="canonical-69ab633e59cb7fbf31f5070a52c1f02b98cfb4853f7ff659c4970470c73d8a3d"></a>

<a id="canonical-6ab4be149674f58000e5ca4593f7de67779cab4a98ac2c2365544c6c81fa5e4c"></a>

## certified_hardware_type property — Property reference / c22b4a524092 / 5

Type: `"string"`. Computed.

\[Enum: VOLTMESH|VOLTSTACK\_COMBO|CLOUD\_MARKET\_PLACE\] Different type of certified HW for billing
rate. Possible values are \`VOLTMESH\`, \`VOLTSTACK\_COMBO\`, \`CLOUD\_MARKET\_PLACE\`. Defaults to
\`VOLTMESH\`.

<a id="canonical-33cfc4da8e0923d72699b1db5c4fb31a02a628c672d62b7885d1ddaab97a3aa2"></a>

<a id="canonical-5663d9824f4cfa931597251b665c75d88f1498cb2f656c5f2d1642188b28398d"></a>

## description property — Property reference / c22b4a524092 / 6

Type: `"string"`. Computed.

Description.

- [devices](data-sources--certified_hardware--reference--group-001.md#canonical-7baeee33e4cc827f25c28ea2cf048f2e5a55ee70814751ca0d455142d3b53c6d): complete subsection reference.

<a id="canonical-35bb162bbbe4ec187239f0a8f478d3e2da5edcbb2a6b55dc2e354f60fa8cc73e"></a>

<a id="canonical-7e853e4b33d1a4b0d4c6b578bacc2965cb4de4aacb3e3b19b2ab19933f6ee015"></a>

## id property — Property reference / c22b4a524092 / 7

Type: `"string"`. Computed.

Unique identifier.

- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68): complete subsection reference.

- [internal_usb_device_rule](data-sources--certified_hardware--reference--group-001.md#canonical-c1df81866f89c2e466059b8c6e253fbf26f763f348e456896453d33dde5d5b04): complete subsection reference.

<a id="canonical-465c586364d09cb5f89eb53168246c7c414799c37d37772a92b1eb45f23445cc"></a>

<a id="canonical-8b5e8394bf220286d56dc3227fce9aaaca3193cbad63ae9296850d3cddac1dde"></a>

## labels property — Property reference / c22b4a524092 / 8

Type: `["map", "string"]`. Computed.

Labels.

<a id="canonical-e61c39571a3c97d9582ecf5841c9809b92c7ca7a6b63cac9e3abe582cec5851e"></a>

<a id="canonical-726716ee375f43c6635796699cd902146aa7494ba2a3a30a644b222c9c09b8c2"></a>

## mem_page_number property — Property reference / c22b4a524092 / 9

Type: `"number"`. Computed.

Number of pages allocated in this certified hardware for Hugepages. Each page size is defined above
in 'mem\_page\_size' Total memory reserved for Hugepages is 'mem\_page\_size \* mem\_page\_number'.

<a id="canonical-a124d00edeac4738facadd1b9fd2e9ded2dc32df4e3b9af54dc081d01addb8bf"></a>

<a id="canonical-1689e28bfbd23c4532b59232c31b7b0700716f3ad408ca8227716289c8dec12f"></a>

## mem_page_size property — Property reference / c22b4a524092 / 10

Type: `"string"`. Computed.

\[Enum:
HARDWARE\_MEM\_PAGE\_SIZE\_INVALID|HARDWARE\_MEM\_PAGE\_SIZE\_4KB|HARDWARE\_MEM\_PAGE\_SIZE\_2MB|HARDWARE\_MEM\_PAGE\_SIZE\_1GB\]
Memory for packets buffers etc are allocated in blocks of pages. Size of each memory page is defined
here. Invalid Page size Page size of 4KB Page size of 2MB Page size of 1GB. Possible values are
\`HARDWARE\_MEM\_PAGE\_SIZE\_INVALID\`, \`HARDWARE\_MEM\_PAGE\_SIZE\_4KB\`,
\`HARDWARE\_MEM\_PAGE\_SIZE\_2MB\`, \`HARDWARE\_MEM\_PAGE\_SIZE\_1GB\`. Defaults to
\`HARDWARE\_MEM\_PAGE\_SIZE\_INVALID\`.

<a id="canonical-d5a493ef97baa01c54811f9f0275291605e3e229ecfeb7474ff87a3c70b72533"></a>

<a id="canonical-78a3b88ab7718dc764cec422a6e69cfdf5e6e30d489f86c93db3dbd3bfb9db3e"></a>

## name property — Property reference / c22b4a524092 / 11

Type: `"string"`. Required.

Name of the CertifiedHardware to look up.

<a id="canonical-d152ddaaeeb5b019087cd8f6d81a4f26fc5df8bd2fa0408acf174e7972c86b50"></a>

<a id="canonical-135b83550edaf71f209b16da332183a572955cec351473a385b508a577332914"></a>

## namespace property — Property reference / c22b4a524092 / 12

Type: `"string"`. Required.

Namespace of the CertifiedHardware.

- [numa_mem](data-sources--certified_hardware--reference--group-001.md#canonical-bcbeeb01b8644d593914c435c9f52c801c39b5740fd9ea3ec09fe5e115c7f893): complete subsection reference.

<a id="canonical-d76ba09e071b2e765f07b7768819eeae1a29c2a8e342de3a73a19cc6b044dff2"></a>

<a id="canonical-409929e95be951e0bfdb84616b534673553b162ffc758e5f17bf65ae17ec1fe5"></a>

## numa_nodes property — Property reference / c22b4a524092 / 13

Type: `"number"`. Computed.

The number of host NUMA nodes used in certified hardware.

- [vendor_model_list](data-sources--certified_hardware--reference--group-001.md#canonical-64519123e849fbc61f12441e0a23a510716d560f280e463dc8d63f16132a1f4d): complete subsection reference.

<a id="canonical-9624ac9ad11c50d6734d8821fb9c349a539d87daf4b0830196eadaebc3920d75"></a>

## All schema paths — Property reference / c22b4a524092 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--certified_hardware--reference--group-001.md#canonical-4d08f3e105e21129a0f84c48563515ab3d62c207500e5ac91c4bab01afaf896c) |
| `certified_hardware_type` | [certified_hardware_type](data-sources--certified_hardware--reference--group-001.md#canonical-69ab633e59cb7fbf31f5070a52c1f02b98cfb4853f7ff659c4970470c73d8a3d) |
| `description` | [description](data-sources--certified_hardware--reference--group-001.md#canonical-33cfc4da8e0923d72699b1db5c4fb31a02a628c672d62b7885d1ddaab97a3aa2) |
| `devices` | [devices](data-sources--certified_hardware--reference--group-001.md#canonical-6215d1286b2bf797eaa214adde91ee3bdb7b5cad13737f10ddff93c4de357c5b) |
| `devices.device_list` | [devices.device_list](data-sources--certified_hardware--reference--group-001.md#canonical-e07b8fb590a6b09344ed33cc7b2ba99acb0d4b231ac79f2b1f34d8dd4f58e5ff) |
| `devices.max_unit` | [devices.max_unit](data-sources--certified_hardware--reference--group-001.md#canonical-a626a9fcc73ddad23f1f88850347b9148e24b296a6244b257de87d51d28c19d0) |
| `devices.min_unit` | [devices.min_unit](data-sources--certified_hardware--reference--group-001.md#canonical-82fe0457244ca936a3db7dd84da41b4b389133a35ecafe22289318416b2ea06a) |
| `devices.name` | [devices.name](data-sources--certified_hardware--reference--group-001.md#canonical-54ef46599a0d7a8cde39a01f6e3d62ac2c3a9e6f2633d6be211b3a5a52641479) |
| `devices.type` | [devices.type](data-sources--certified_hardware--reference--group-001.md#canonical-7bce2bc7af67d85cdc894056920bda3da3b91b9bd332202efc5a951ddf774eaa) |
| `devices.use` | [devices.use](data-sources--certified_hardware--reference--group-001.md#canonical-a6d61c1109ccf7361829e443ab650c1551467dbda9f51140622172c2a6b772e3) |
| `id` | [id](data-sources--certified_hardware--reference--group-001.md#canonical-35bb162bbbe4ec187239f0a8f478d3e2da5edcbb2a6b55dc2e354f60fa8cc73e) |
| `image_list` | [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-97152f757484b900c37cd1225eead55199886749af0a5ba68d11d96a6e6c8860) |
| `image_list.aws` | [image_list.aws](data-sources--certified_hardware--reference--group-001.md#canonical-659ea93b0c8cea8d6d31e2c615eb14b0364fbc449a62f9beab9c0d7625c1231c) |
| `image_list.aws.image_id` | [image_list.aws.image_id](data-sources--certified_hardware--reference--group-001.md#canonical-075e3b678057f956d9994c343cbca0fb3f572808759d0a4e3626955eee471f4b) |
| `image_list.aws.image_id.image_id` | [image_list.aws.image_id.image_id](data-sources--certified_hardware--reference--group-001.md#canonical-0b118ce1841c8672d3a1e4561712ace727649ca1c5ef5a9c8399189c8c655abd) |
| `image_list.aws.image_id.region` | [image_list.aws.image_id.region](data-sources--certified_hardware--reference--group-001.md#canonical-e71d6bb810de1f25ea29684703e103636b919e1f636f3a29870803b081aed755) |
| `image_list.azure` | [image_list.azure](data-sources--certified_hardware--reference--group-001.md#canonical-7b67940caaa4b793c97b36192436c6b8f78d4a766a671f4c66b53278cf91381c) |
| `image_list.azure.image_id` | [image_list.azure.image_id](data-sources--certified_hardware--reference--group-001.md#canonical-31393b10813fb744a2fa17d63340f89af8697c9da2912e9e1e270efc49966dae) |
| `image_list.azure.image_id.image_id` | [image_list.azure.image_id.image_id](data-sources--certified_hardware--reference--group-001.md#canonical-ab983bcd941742d120008055ac280511b2a718765f8552b1d0d042f762deffe2) |
| `image_list.azure.marketplace` | [image_list.azure.marketplace](data-sources--certified_hardware--reference--group-001.md#canonical-be75a72ecee2c6a6bfcbb25898f6e13d235dc92d922387cfd8b789dd4105ac83) |
| `image_list.azure.marketplace.name` | [image_list.azure.marketplace.name](data-sources--certified_hardware--reference--group-001.md#canonical-f3b07ff1333b23d62277d5dd9c44c9d5429a15acf697da5c5d950dbf6c95053a) |
| `image_list.azure.marketplace.offer` | [image_list.azure.marketplace.offer](data-sources--certified_hardware--reference--group-001.md#canonical-8a4c1e8de0e23df67276465bb29e538017f0e317f9743773833a586661e15518) |
| `image_list.azure.marketplace.publisher` | [image_list.azure.marketplace.publisher](data-sources--certified_hardware--reference--group-001.md#canonical-bbe8de719f3d63083de89c6688e1bdaedac8765f0a5167ec33ffc904fce72aad) |
| `image_list.azure.marketplace.sku` | [image_list.azure.marketplace.sku](data-sources--certified_hardware--reference--group-001.md#canonical-02957541b165a13d4aa7b35bd82836a2882d77e3a85258b4f1c73e3e901fc11f) |
| `image_list.azure.marketplace.version` | [image_list.azure.marketplace.version](data-sources--certified_hardware--reference--group-001.md#canonical-2bd5435c9c38fe25d439353a4984263fc0e29447e7a53d1815b60dff4d313d85) |
| `image_list.gcp` | [image_list.gcp](data-sources--certified_hardware--reference--group-001.md#canonical-e8e464d280054295151f6bc1bac4d6145d7d61282ded6a2fd2865883350e2e26) |
| `image_list.gcp.image_id` | [image_list.gcp.image_id](data-sources--certified_hardware--reference--group-001.md#canonical-7a8abefd4f48deb7069dca9c2bac3b58107376674f95f8309f538fa38eb4d472) |
| `image_list.gcp.image_id.image_id` | [image_list.gcp.image_id.image_id](data-sources--certified_hardware--reference--group-001.md#canonical-df5d4c4c1dd176525394e3eaf07e067227075745d8b75c032cedfda40ab60ec5) |
| `image_list.name` | [image_list.name](data-sources--certified_hardware--reference--group-001.md#canonical-a9a16b8c4a60ed80440595a77de6f3575ca19f1c091343f6f7d75b9b3b7dd49a) |
| `image_list.provider_ref` | [image_list.provider_ref](data-sources--certified_hardware--reference--group-001.md#canonical-1451b2906ab3f446831005c91df89c023a5fe5d310a695af4ca8b9ec6afd06a2) |
| `internal_usb_device_rule` | [internal_usb_device_rule](data-sources--certified_hardware--reference--group-001.md#canonical-4ff16b98172b0c4d1ed4002782f2cbd09dbf7675e0bd0d5d353975a4f24eab83) |
| `internal_usb_device_rule.b_device_class` | [internal_usb_device_rule.b_device_class](data-sources--certified_hardware--reference--group-001.md#canonical-f647bc0fbcc121989d13f89fe250ccb2e67622168fda6b14039842914ffda9cb) |
| `internal_usb_device_rule.b_device_protocol` | [internal_usb_device_rule.b_device_protocol](data-sources--certified_hardware--reference--group-001.md#canonical-7be8ceafc553ed2de656381696d7b53ecba2fafd4dade34ed8842b2f8cdc4cf0) |
| `internal_usb_device_rule.b_device_sub_class` | [internal_usb_device_rule.b_device_sub_class](data-sources--certified_hardware--reference--group-001.md#canonical-7bd40ad9f863332fc9f958176216fcc085b789d9cc149037f559943473f51aba) |
| `internal_usb_device_rule.i_serial` | [internal_usb_device_rule.i_serial](data-sources--certified_hardware--reference--group-001.md#canonical-4d0bacd2cb82d1ec950219da3c08b3c74b3eb08659f25c46acc7a11960342970) |
| `internal_usb_device_rule.id_product` | [internal_usb_device_rule.id_product](data-sources--certified_hardware--reference--group-001.md#canonical-d003c1105033e4b1aa814bdb0e739585780f8e4977a1c401fc00b241868a1444) |
| `internal_usb_device_rule.id_vendor` | [internal_usb_device_rule.id_vendor](data-sources--certified_hardware--reference--group-001.md#canonical-168afb7c688492c411b878087fe734420c1ac2dc312c8b9f266c607111b78512) |
| `labels` | [labels](data-sources--certified_hardware--reference--group-001.md#canonical-465c586364d09cb5f89eb53168246c7c414799c37d37772a92b1eb45f23445cc) |
| `mem_page_number` | [mem_page_number](data-sources--certified_hardware--reference--group-001.md#canonical-e61c39571a3c97d9582ecf5841c9809b92c7ca7a6b63cac9e3abe582cec5851e) |
| `mem_page_size` | [mem_page_size](data-sources--certified_hardware--reference--group-001.md#canonical-a124d00edeac4738facadd1b9fd2e9ded2dc32df4e3b9af54dc081d01addb8bf) |
| `name` | [name](data-sources--certified_hardware--reference--group-001.md#canonical-d5a493ef97baa01c54811f9f0275291605e3e229ecfeb7474ff87a3c70b72533) |
| `namespace` | [namespace](data-sources--certified_hardware--reference--group-001.md#canonical-d152ddaaeeb5b019087cd8f6d81a4f26fc5df8bd2fa0408acf174e7972c86b50) |
| `numa_mem` | [numa_mem](data-sources--certified_hardware--reference--group-001.md#canonical-34c69447c6622c4fc7466b2d309d668391f08a792c790ba6ddccf4b829ad7b23) |
| `numa_mem.memory` | [numa_mem.memory](data-sources--certified_hardware--reference--group-001.md#canonical-d0bb4fcb46a8fad0a36d2a376de1b889746e7f5c8189e5f6680b351b2d91db4c) |
| `numa_mem.node` | [numa_mem.node](data-sources--certified_hardware--reference--group-001.md#canonical-5a999149527e25c95ab281fe6b7b87d5c0377e37c598a241861b47207bc5cf0e) |
| `numa_nodes` | [numa_nodes](data-sources--certified_hardware--reference--group-001.md#canonical-d76ba09e071b2e765f07b7768819eeae1a29c2a8e342de3a73a19cc6b044dff2) |
| `vendor_model_list` | [vendor_model_list](data-sources--certified_hardware--reference--group-001.md#canonical-13dc07330e648eca7218d58d42ac422f4f1ab7262910fbf41b1eb3d553947a2c) |
| `vendor_model_list.model` | [vendor_model_list.model](data-sources--certified_hardware--reference--group-001.md#canonical-bc2f307329047f45a9f35fdae12b94898a30db5c0dc02d39c066ac34d608b449) |
| `vendor_model_list.vendor` | [vendor_model_list.vendor](data-sources--certified_hardware--reference--group-001.md#canonical-318277bd5075e0a985c701d7dc6896b83b715b9fafb5f50c050c0f89ad4a66c5) |

<a id="canonical-9164de2f1f95c66e3e368de5653f663edf8c22e526b6567fc84f11e6e7a81a4f"></a>

## Next pages — Property reference / c22b4a524092 / 15

- [devices](data-sources--certified_hardware--reference--group-001.md#canonical-7baeee33e4cc827f25c28ea2cf048f2e5a55ee70814751ca0d455142d3b53c6d)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- [internal_usb_device_rule](data-sources--certified_hardware--reference--group-001.md#canonical-c1df81866f89c2e466059b8c6e253fbf26f763f348e456896453d33dde5d5b04)
- [numa_mem](data-sources--certified_hardware--reference--group-001.md#canonical-bcbeeb01b8644d593914c435c9f52c801c39b5740fd9ea3ec09fe5e115c7f893)
- [vendor_model_list](data-sources--certified_hardware--reference--group-001.md#canonical-64519123e849fbc61f12441e0a23a510716d560f280e463dc8d63f16132a1f4d)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-7baeee33e4cc827f25c28ea2cf048f2e5a55ee70814751ca0d455142d3b53c6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-616d4d650f178d6b3418f17e637c80ea2d057f1c4008989dac256815a90d2c5f"></a>

## devices — devices / 2b4ea207163d / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- devices

<a id="canonical-6215d1286b2bf797eaa214adde91ee3bdb7b5cad13737f10ddff93c4de357c5b"></a>

Type: `"list"`. Computed.

List of supported devices in this model.

<a id="canonical-78250ea2ffeceb430ea6cea6cc08b89c2bb596e43dc71126bdb6463fa889e495"></a>

## Direct properties — devices / 2b4ea207163d / 3

<a id="canonical-e07b8fb590a6b09344ed33cc7b2ba99acb0d4b231ac79f2b1f34d8dd4f58e5ff"></a>

<a id="canonical-ced9e11c3f11c0d71112e32812a60cace6709f9285f6fe83f60146c13139498f"></a>

## device_list property — devices / 2b4ea207163d / 4

Type: `["list", "string"]`. Computed.

In case of logical boot strap devices like LACP Link aggregation or RAID.

<a id="canonical-a626a9fcc73ddad23f1f88850347b9148e24b296a6244b257de87d51d28c19d0"></a>

<a id="canonical-6c757b9fd671c63c65319092166323a20f97e37f49f21a696c39e08dfcb3f5b0"></a>

## max_unit property — devices / 2b4ea207163d / 5

Type: `"number"`. Computed.

Last unit number of the device supported in this certified hardware.

<a id="canonical-82fe0457244ca936a3db7dd84da41b4b389133a35ecafe22289318416b2ea06a"></a>

<a id="canonical-5035d02c3028f92fc50c5d5386149341862fe23d0f1b3fa0d3648d3ad7f521c1"></a>

## min_unit property — devices / 2b4ea207163d / 6

Type: `"number"`. Computed.

First unit number of the device supported in this certified hardware.

<a id="canonical-54ef46599a0d7a8cde39a01f6e3d62ac2c3a9e6f2633d6be211b3a5a52641479"></a>

<a id="canonical-163cd499a1a690925195de88eb7823330d126d037419b427938b6f51572575eb"></a>

## name property — devices / 2b4ea207163d / 7

Type: `"string"`. Computed.

Device. Name of the device.

<a id="canonical-7bce2bc7af67d85cdc894056920bda3da3b91b9bd332202efc5a951ddf774eaa"></a>

<a id="canonical-1aabb4afc1a0372aebbb219d9fa386e7c8b7e6eb9a244adeaf84fb1dffe2bf06"></a>

## type property — devices / 2b4ea207163d / 8

Type: `"string"`. Computed.

\[Enum:
HARDWARE\_DEVICE\_INVALID|HARDWARE\_DEVICE\_ETHERNET|HARDWARE\_DEVICE\_VIRTIO|HARDWARE\_DEVICE\_TUNTAP|HARDWARE\_DEVICE\_BOND|HARDWARE\_DEVICE\_EXTERNAL\_ISCSI\_STORTAGE|HARDWARE\_DEVICE\_NVIDIA\_GPU\]
Different type of devices supported Invalid device or device that's not supported Ethernet device
VIRTIO device TUNTAP device LACP based bond interface External iSCSI devices supported Nvidia GPU
device used for machine learning. Possible values are \`HARDWARE\_DEVICE\_INVALID\`,
\`HARDWARE\_DEVICE\_ETHERNET\`, \`HARDWARE\_DEVICE\_VIRTIO\`, \`HARDWARE\_DEVICE\_TUNTAP\`,
\`HARDWARE\_DEVICE\_BOND\`, \`HARDWARE\_DEVICE\_EXTERNAL\_ISCSI\_STORTAGE\`,
\`HARDWARE\_DEVICE\_NVIDIA\_GPU\`. Defaults to \`HARDWARE\_DEVICE\_INVALID\`.

<a id="canonical-a6d61c1109ccf7361829e443ab650c1551467dbda9f51140622172c2a6b772e3"></a>

<a id="canonical-a4ecc79075cc9293ff0112f3a556d5c4f325b538bfcb41f88af4e84fe349b5c6"></a>

## use property — devices / 2b4ea207163d / 9

Type: `"string"`. Computed.

\[Enum:
HARDWARE\_DEVICE\_USE\_REGULAR|HARDWARE\_DEVICE\_USE\_INTERNAL|HARDWARE\_NETWORK\_DEVICE\_USE\_REGULAR|HARDWARE\_NETWORK\_DEVICE\_USE\_INTERNAL|HARDWARE\_NETWORK\_DEVICE\_USE\_MANAGEMENT|HARDWARE\_NETWORK\_DEVICE\_USE\_OUTSIDE|HARDWARE\_NETWORK\_DEVICE\_USE\_INSIDE|HARDWARE\_NETWORK\_DEVICE\_USE\_OUTSIDE\_LAG|HARDWARE\_NETWORK\_DEVICE\_USE\_INSIDE\_LAG|HARDWARE\_NETWORK\_DEVICE\_USE\_LAG\_MEMBER|HARDWARE\_NETWORK\_DEVICE\_USE\_STORAGE|HARDWARE\_NETWORK\_DEVICE\_USE\_FALLBACK\_MANAGEMENT\]
Defines how the device instance must be used If the device is owned by F5 Distributed Cloud
software, it is available for users to configure as required Device reserved for internal use by F5
Distributed Cloud Node If the Network device is owned by VER, it is available for users to configure
as.. Possible values are \`HARDWARE\_DEVICE\_USE\_REGULAR\`, \`HARDWARE\_DEVICE\_USE\_INTERNAL\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_REGULAR\`, \`HARDWARE\_NETWORK\_DEVICE\_USE\_INTERNAL\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_MANAGEMENT\`, \`HARDWARE\_NETWORK\_DEVICE\_USE\_OUTSIDE\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_INSIDE\`, \`HARDWARE\_NETWORK\_DEVICE\_USE\_OUTSIDE\_LAG\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_INSIDE\_LAG\`, \`HARDWARE\_NETWORK\_DEVICE\_USE\_LAG\_MEMBER\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_STORAGE\`,
\`HARDWARE\_NETWORK\_DEVICE\_USE\_FALLBACK\_MANAGEMENT\`. Defaults to
\`HARDWARE\_DEVICE\_USE\_REGULAR\`.

<a id="canonical-247e41a963a1e884abf8fd46b338ed7908cabad4e1090703774d83d720a19147"></a>

## Next pages — devices / 2b4ea207163d / 10

- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-244344446088721c6c6de99a9be264812448ab9f6075548823e52528c19dcb7e"></a>

## image_list — image_list / 1d324b72ff87 / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- image_list

<a id="canonical-97152f757484b900c37cd1225eead55199886749af0a5ba68d11d96a6e6c8860"></a>

Type: `"list"`. Computed.

List of image names with providers for this certified hardware, e.g. AWS ami-0f99d090261d2acd5.

<a id="canonical-3b8fe5736f3394e49711fd4cf0b37898b29d5d0ec01364dc9268a406e279a10b"></a>

## Direct properties — image_list / 1d324b72ff87 / 3

- [aws](data-sources--certified_hardware--reference--group-001.md#canonical-5bf2a2ab44350101d18716a0acb19b55744c0f73717128836926626bd96bb36e): complete subsection reference.

- [azure](data-sources--certified_hardware--reference--group-001.md#canonical-fbe93ad3745a72536da6e8032314cc4ab003b96dba1bb2931b0572ddcf766839): complete subsection reference.

- [gcp](data-sources--certified_hardware--reference--group-001.md#canonical-68b29cb40c6337ed5d97842ee83dfa9f492c8c37cb7c0d106dae375eb11539f4): complete subsection reference.

<a id="canonical-a9a16b8c4a60ed80440595a77de6f3575ca19f1c091343f6f7d75b9b3b7dd49a"></a>

<a id="canonical-0b8de5e643f69a569f84c6c806fcaa36287ba2083bf96d9563750d3edc7327c2"></a>

## name property — image_list / 1d324b72ff87 / 4

Type: `"string"`. Computed.

Name. Image name to use for this hardware.

<a id="canonical-1451b2906ab3f446831005c91df89c023a5fe5d310a695af4ca8b9ec6afd06a2"></a>

<a id="canonical-19be280d565b3492d4583f60a98a867716fc68148d51a42d2e1e2a69831c51f5"></a>

## provider_ref property — image_list / 1d324b72ff87 / 5

Type: `"string"`. Computed.

Image provider F5 Distributed Cloud, Cloud provider like AWS or Azure.

<a id="canonical-6d6f9707ac8bbfbbd8e2800353668c1edaa6dbd4f41e0e95e54ff646a837cd0e"></a>

## Next pages — image_list / 1d324b72ff87 / 6

- [image_list.aws](data-sources--certified_hardware--reference--group-001.md#canonical-5bf2a2ab44350101d18716a0acb19b55744c0f73717128836926626bd96bb36e)
- [image_list.azure](data-sources--certified_hardware--reference--group-001.md#canonical-fbe93ad3745a72536da6e8032314cc4ab003b96dba1bb2931b0572ddcf766839)
- [image_list.gcp](data-sources--certified_hardware--reference--group-001.md#canonical-68b29cb40c6337ed5d97842ee83dfa9f492c8c37cb7c0d106dae375eb11539f4)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-5bf2a2ab44350101d18716a0acb19b55744c0f73717128836926626bd96bb36e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0195bad3464bd3d9f3d616f41978b9bdaa442b50a2f1be4a67c1b81738ac101c"></a>

## image_list.aws — image_list.aws / de915ea75214 / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- image_list.aws

<a id="canonical-659ea93b0c8cea8d6d31e2c615eb14b0364fbc449a62f9beab9c0d7625c1231c"></a>

Type: `"single"`. Computed.

AWS. AWS specific information.

<a id="canonical-b82816367ab6dd2965aa687cd30f37b0ce0a12c5c2fcd7ab97a9ae4a8acf7e04"></a>

## Direct properties — image_list.aws / de915ea75214 / 3

- [image_id](data-sources--certified_hardware--reference--group-001.md#canonical-72fe32b425befc46c1919ac5638bf8665dab139284f935068a5b4bb59bf12072): complete subsection reference.

<a id="canonical-976c3a50d6dfaaa8ae51fe5281810d96c1ac337448f7d0b25e974ebfdf7e1731"></a>

## Next pages — image_list.aws / de915ea75214 / 4

- [image_list.aws.image_id](data-sources--certified_hardware--reference--group-001.md#canonical-72fe32b425befc46c1919ac5638bf8665dab139284f935068a5b4bb59bf12072)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-72fe32b425befc46c1919ac5638bf8665dab139284f935068a5b4bb59bf12072"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f460da240dc787ea815f01f65fdaee4b9e1789a0b5b384acbc91436c5d374a2a"></a>

## image_list.aws.image_id — image_list.aws.image_id / b5a8b302c662 / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- [image_list.aws](data-sources--certified_hardware--reference--group-001.md#canonical-5bf2a2ab44350101d18716a0acb19b55744c0f73717128836926626bd96bb36e)
- image_list.aws.image_id

<a id="canonical-075e3b678057f956d9994c343cbca0fb3f572808759d0a4e3626955eee471f4b"></a>

Type: `"single"`. Computed.

Configuration for image\_id.

<a id="canonical-14520145ce77a74ef8a6f0b36535d4aa64dccec7d6670677d3865797f0231794"></a>

## Direct properties — image_list.aws.image_id / b5a8b302c662 / 3

<a id="canonical-0b118ce1841c8672d3a1e4561712ace727649ca1c5ef5a9c8399189c8c655abd"></a>

<a id="canonical-1c803b8e6b1d9b12caaf04c2303443c5da4dcda1c9ba1d63b215b81ee21179bd"></a>

## image_id property — image_list.aws.image_id / b5a8b302c662 / 4

Type: `"string"`. Computed.

AWS ami image name. AWS ami image.

<a id="canonical-e71d6bb810de1f25ea29684703e103636b919e1f636f3a29870803b081aed755"></a>

<a id="canonical-28f49d73e69524987f11f858e02d3a42973652c789769dec741d18291fa3d691"></a>

## region property — image_list.aws.image_id / b5a8b302c662 / 5

Type: `"string"`. Computed.

AWS ami image region. AWS ami image region.

<a id="canonical-3b84e56e3012d706b776e849cba494e8804d2983592b2c513d26432d08a10cf6"></a>

## Next pages — image_list.aws.image_id / b5a8b302c662 / 6

- [image_list.aws](data-sources--certified_hardware--reference--group-001.md#canonical-5bf2a2ab44350101d18716a0acb19b55744c0f73717128836926626bd96bb36e)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-fbe93ad3745a72536da6e8032314cc4ab003b96dba1bb2931b0572ddcf766839"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eed6dc4cb5b82135fdb4cc007c382c6ee5f8c2333940d3514eea64987ee9dd7c"></a>

## image_list.azure — image_list.azure / a1d40ba7c21a / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- image_list.azure

<a id="canonical-7b67940caaa4b793c97b36192436c6b8f78d4a766a671f4c66b53278cf91381c"></a>

Type: `"single"`. Computed.

Azure. Azure specific information.

<a id="canonical-baf3a2d2eb0353813d217d9fb178bd4354d732dc2525a556e0862fd494930626"></a>

## Direct properties — image_list.azure / a1d40ba7c21a / 3

- [image_id](data-sources--certified_hardware--reference--group-001.md#canonical-33d37ec4d48ace1e54686790ca54b17bebb8be15be02bb773310be7972e08e0b): complete subsection reference.

- [marketplace](data-sources--certified_hardware--reference--group-001.md#canonical-13617fc919fc6e08edf2dc1de01eb73c38cbded805f56bc0366da9d51ea3a73d): complete subsection reference.

<a id="canonical-569a0d9101a01cd77a2ea399886e3fa3875138c0f320acf9f5b1325373befd25"></a>

## Next pages — image_list.azure / a1d40ba7c21a / 4

- [image_list.azure.image_id](data-sources--certified_hardware--reference--group-001.md#canonical-33d37ec4d48ace1e54686790ca54b17bebb8be15be02bb773310be7972e08e0b)
- [image_list.azure.marketplace](data-sources--certified_hardware--reference--group-001.md#canonical-13617fc919fc6e08edf2dc1de01eb73c38cbded805f56bc0366da9d51ea3a73d)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-33d37ec4d48ace1e54686790ca54b17bebb8be15be02bb773310be7972e08e0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8a73bc297eb88c55a953c22e60dfa5a8e236f3b8e28ea1c01209c7c34dbd78f"></a>

## image_list.azure.image_id — image_list.azure.image_id / 6435d2261c87 / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- [image_list.azure](data-sources--certified_hardware--reference--group-001.md#canonical-fbe93ad3745a72536da6e8032314cc4ab003b96dba1bb2931b0572ddcf766839)
- image_list.azure.image_id

<a id="canonical-31393b10813fb744a2fa17d63340f89af8697c9da2912e9e1e270efc49966dae"></a>

Type: `"single"`. Computed.

Configuration for image\_id.

<a id="canonical-ad87e5634a5d4b524225620297ba0f0347eb07bd98f467c1949b64ef15b37ce8"></a>

## Direct properties — image_list.azure.image_id / 6435d2261c87 / 3

<a id="canonical-ab983bcd941742d120008055ac280511b2a718765f8552b1d0d042f762deffe2"></a>

<a id="canonical-034a764fc576bcd067be073a1b84b7f53ee02474823cf3843686ecabf5a12c1f"></a>

## image_id property — image_list.azure.image_id / 6435d2261c87 / 4

Type: `"string"`. Computed.

Azure image ID info. Azure image ID.

<a id="canonical-3fffd95a2dbbbcfa1bde14afed293f041537b9886c543e0391b88eb11e51ba91"></a>

## Next pages — image_list.azure.image_id / 6435d2261c87 / 5

- [image_list.azure](data-sources--certified_hardware--reference--group-001.md#canonical-fbe93ad3745a72536da6e8032314cc4ab003b96dba1bb2931b0572ddcf766839)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-13617fc919fc6e08edf2dc1de01eb73c38cbded805f56bc0366da9d51ea3a73d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b47a78817a0717dc244302e6d85ac9b63a8852fb2cebeb487b246a6d0a1b1f8"></a>

## image_list.azure.marketplace — image_list.azure.marketplace / b4d8785ad7eb / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- [image_list.azure](data-sources--certified_hardware--reference--group-001.md#canonical-fbe93ad3745a72536da6e8032314cc4ab003b96dba1bb2931b0572ddcf766839)
- image_list.azure.marketplace

<a id="canonical-be75a72ecee2c6a6bfcbb25898f6e13d235dc92d922387cfd8b789dd4105ac83"></a>

Type: `"single"`. Computed.

Configuration parameter for marketplace.

<a id="canonical-f499f79ae4a34d540127412e8951cac482c2f081fcc1c1581f5132c8329e3546"></a>

## Direct properties — image_list.azure.marketplace / b4d8785ad7eb / 3

<a id="canonical-f3b07ff1333b23d62277d5dd9c44c9d5429a15acf697da5c5d950dbf6c95053a"></a>

<a id="canonical-abdda16a5dcc659550a041e6a5874bbdd92d446997919fc2dc6b58075a7e8ba4"></a>

## name property — image_list.azure.marketplace / b4d8785ad7eb / 4

Type: `"string"`. Computed.

Azure Marketplace Name. Azure Marketplace Name.

<a id="canonical-8a4c1e8de0e23df67276465bb29e538017f0e317f9743773833a586661e15518"></a>

<a id="canonical-0ce731727c4a7d6a3179681cb4d1db8a1ca627ef50b9dd495e2a588ed8a68071"></a>

## offer property — image_list.azure.marketplace / b4d8785ad7eb / 5

Type: `"string"`. Computed.

Azure Marketplace offer. Azure Marketplace offer.

<a id="canonical-bbe8de719f3d63083de89c6688e1bdaedac8765f0a5167ec33ffc904fce72aad"></a>

<a id="canonical-ae0e0428c89a61e09ff4479a1c580f1aa2f4dcc6d1b0f33b73ac1a85a8318f4f"></a>

## publisher property — image_list.azure.marketplace / b4d8785ad7eb / 6

Type: `"string"`. Computed.

Azure Marketplace Publisher. Azure Marketplace Publisher.

<a id="canonical-02957541b165a13d4aa7b35bd82836a2882d77e3a85258b4f1c73e3e901fc11f"></a>

<a id="canonical-efc07dbeb291c819a73257c9b90049a01964b769577b721e16ece588148dc10a"></a>

## sku property — image_list.azure.marketplace / b4d8785ad7eb / 7

Type: `"string"`. Computed.

Azure Marketplace SKU. Azure Marketplace SKU.

<a id="canonical-2bd5435c9c38fe25d439353a4984263fc0e29447e7a53d1815b60dff4d313d85"></a>

<a id="canonical-ff664283ce38db3af73df766ea7e1442ede90908a7f76aed822d65707dc8ebb1"></a>

## version property — image_list.azure.marketplace / b4d8785ad7eb / 8

Type: `"string"`. Computed.

Azure Marketplace Version. Azure Marketplace Version.

<a id="canonical-a22b17dfb6252d5e4365d5f2a68e096f10d85ebd9c3157818c0fa7b4590ad4c1"></a>

## Next pages — image_list.azure.marketplace / b4d8785ad7eb / 9

- [image_list.azure](data-sources--certified_hardware--reference--group-001.md#canonical-fbe93ad3745a72536da6e8032314cc4ab003b96dba1bb2931b0572ddcf766839)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-68b29cb40c6337ed5d97842ee83dfa9f492c8c37cb7c0d106dae375eb11539f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59a91692829ef7b5db4c1aedeac9c42b75a4c87074179a952d9214a7db1ac0ed"></a>

## image_list.gcp — image_list.gcp / 4bea4df63329 / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- image_list.gcp

<a id="canonical-e8e464d280054295151f6bc1bac4d6145d7d61282ded6a2fd2865883350e2e26"></a>

Type: `"single"`. Computed.

GCP. GCP specific information.

<a id="canonical-e346a4a2d682f826e421dd4ae6299bbcda7f9c710bcd28c999ce710c4fc44c4b"></a>

## Direct properties — image_list.gcp / 4bea4df63329 / 3

- [image_id](data-sources--certified_hardware--reference--group-001.md#canonical-cf8880e4c48b04d945de8e0ba8be1a530f0cacec72ccbe7eb8dd69193a523a34): complete subsection reference.

<a id="canonical-e0782196de631de574bcb00a7c2eee1da4404429d1514ba83886deaf9d3a75d8"></a>

## Next pages — image_list.gcp / 4bea4df63329 / 4

- [image_list.gcp.image_id](data-sources--certified_hardware--reference--group-001.md#canonical-cf8880e4c48b04d945de8e0ba8be1a530f0cacec72ccbe7eb8dd69193a523a34)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-cf8880e4c48b04d945de8e0ba8be1a530f0cacec72ccbe7eb8dd69193a523a34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48287707d3cf7ab8cf48b9ed41fb6e89413528db9635d1705560e5edce7bfe66"></a>

## image_list.gcp.image_id — image_list.gcp.image_id / b0d70b85d16e / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [image_list](data-sources--certified_hardware--reference--group-001.md#canonical-f186d891158e4bf4c4e5365f5476155bf448d25819982cf76886e337b1701c68)
- [image_list.gcp](data-sources--certified_hardware--reference--group-001.md#canonical-68b29cb40c6337ed5d97842ee83dfa9f492c8c37cb7c0d106dae375eb11539f4)
- image_list.gcp.image_id

<a id="canonical-7a8abefd4f48deb7069dca9c2bac3b58107376674f95f8309f538fa38eb4d472"></a>

Type: `"single"`. Computed.

Configuration for image\_id.

<a id="canonical-f043420341d435f2eaf5165f98849dfd62762ee4ba6fc943b28fb80c4865654d"></a>

## Direct properties — image_list.gcp.image_id / b0d70b85d16e / 3

<a id="canonical-df5d4c4c1dd176525394e3eaf07e067227075745d8b75c032cedfda40ab60ec5"></a>

<a id="canonical-c1c2427e196c91600e7af2f61e4e8998a884e4ae6fce960df721ad6358df99e7"></a>

## image_id property — image_list.gcp.image_id / b0d70b85d16e / 4

Type: `"string"`. Computed.

GCP image name. GCP image

<a id="canonical-56a07ee2fa73d0b83b54a8062fcb88f0942123ea5827f44b81d75ebb90e7b2d1"></a>

## Next pages — image_list.gcp.image_id / b0d70b85d16e / 5

- [image_list.gcp](data-sources--certified_hardware--reference--group-001.md#canonical-68b29cb40c6337ed5d97842ee83dfa9f492c8c37cb7c0d106dae375eb11539f4)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-c1df81866f89c2e466059b8c6e253fbf26f763f348e456896453d33dde5d5b04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03659e52f5b2e24e93e7172f44680120eb530730f689f130e9bdcce74c0ab201"></a>

## internal_usb_device_rule — internal_usb_device_rule / 5f9e2d1e961c / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- internal_usb_device_rule

<a id="canonical-4ff16b98172b0c4d1ed4002782f2cbd09dbf7675e0bd0d5d353975a4f24eab83"></a>

Type: `"list"`. Computed.

List of internal USB device rules for server.

<a id="canonical-f86d0af9307f3c497d2c082764ae646af3782fce0455953b7e9bcd7586c9e3a1"></a>

## Direct properties — internal_usb_device_rule / 5f9e2d1e961c / 3

<a id="canonical-f647bc0fbcc121989d13f89fe250ccb2e67622168fda6b14039842914ffda9cb"></a>

<a id="canonical-4110dfb20dfb311ff69885ba28409e7af350c45bbaf8473b79e83f83d5011d4d"></a>

## b_device_class property — internal_usb_device_rule / 5f9e2d1e961c / 4

Type: `"string"`. Computed.

Class. The class of this device.

<a id="canonical-7be8ceafc553ed2de656381696d7b53ecba2fafd4dade34ed8842b2f8cdc4cf0"></a>

<a id="canonical-a592220d292a1f008fd42d0d695b9c89ebe057ca57ccb2a5a964b447dc7c993b"></a>

## b_device_protocol property — internal_usb_device_rule / 5f9e2d1e961c / 5

Type: `"string"`. Computed.

The protocol (within the sub-class) of this device.

<a id="canonical-7bd40ad9f863332fc9f958176216fcc085b789d9cc149037f559943473f51aba"></a>

<a id="canonical-d37bf6ff2022a88237217dd8c83fba2352faf0717e86e90d4bbeeab90802382c"></a>

## b_device_sub_class property — internal_usb_device_rule / 5f9e2d1e961c / 6

Type: `"string"`. Computed.

The sub-class (within the class) of this device.

<a id="canonical-4d0bacd2cb82d1ec950219da3c08b3c74b3eb08659f25c46acc7a11960342970"></a>

<a id="canonical-aa5253a996f333eab3ba733c07edadba1b3d0e4149095e8a3fadc9b484dd51d0"></a>

## i_serial property — internal_usb_device_rule / 5f9e2d1e961c / 7

Type: `"string"`. Computed.

Index of Serial Number String Descriptor.

<a id="canonical-d003c1105033e4b1aa814bdb0e739585780f8e4977a1c401fc00b241868a1444"></a>

<a id="canonical-499f846416b7284e8cb98bbb27e4c4b7d71fa22b8a7ef0bc46d4c9fde9b1a747"></a>

## id_product property — internal_usb_device_rule / 5f9e2d1e961c / 8

Type: `"string"`. Computed.

Product ID (Assigned by Manufacturer) in hex.

<a id="canonical-168afb7c688492c411b878087fe734420c1ac2dc312c8b9f266c607111b78512"></a>

<a id="canonical-f8e6c7cdb98c5ce99475e8e6cbf1ac2010a48e9d6f837e14618d212b133abf8d"></a>

## id_vendor property — internal_usb_device_rule / 5f9e2d1e961c / 9

Type: `"string"`. Computed.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

<a id="canonical-b8cb1405a2381c6f2e7ee793ba86b2c219dbe57951f4439edbef3711ad1d4977"></a>

## Next pages — internal_usb_device_rule / 5f9e2d1e961c / 10

- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-bcbeeb01b8644d593914c435c9f52c801c39b5740fd9ea3ec09fe5e115c7f893"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd0c92fa1f0bbbeac05aef87f65334dade9dc4e0ea3be943e1e0c20c8e2fb1e5"></a>

## numa_mem — numa_mem / 75e34d6b115c / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- numa_mem

<a id="canonical-34c69447c6622c4fc7466b2d309d668391f08a792c790ba6ddccf4b829ad7b23"></a>

Type: `"list"`. Computed.

List of Numa nodes with the number of MB of instance memory to map to node instance If not
specified, memory is evenly divided among available NUMA nodes.

<a id="canonical-cce98329519f22620be796f8f7e3455d4b7cea5b91416051cb7e531fd5a4176c"></a>

## Direct properties — numa_mem / 75e34d6b115c / 3

<a id="canonical-d0bb4fcb46a8fad0a36d2a376de1b889746e7f5c8189e5f6680b351b2d91db4c"></a>

<a id="canonical-d3b865c59328e67206a29bbb5da6f04f59aca601ef5ba5d2c5e84483bfb786bf"></a>

## memory property — numa_mem / 75e34d6b115c / 4

Type: `"number"`. Computed.

The number of MB of instance memory to map to instance NUMA node N.

<a id="canonical-5a999149527e25c95ab281fe6b7b87d5c0377e37c598a241861b47207bc5cf0e"></a>

<a id="canonical-0405ed4a3463a02e769abfda072946a8a5a5372791a743d16ef768f75005fc6f"></a>

## node property — numa_mem / 75e34d6b115c / 5

Type: `"number"`. Computed.

Node. NUMA node instance with mapped memory.

<a id="canonical-f1599e98628cac8122c02546924eefa68b65580147d3a6566de8eb05415a9829"></a>

## Next pages — numa_mem / 75e34d6b115c / 6

- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-64519123e849fbc61f12441e0a23a510716d560f280e463dc8d63f16132a1f4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45a3a541a6811f9d3f05f5b5e78479a28792bc2cc400e7bc44662949970aab44"></a>

## vendor_model_list — vendor_model_list / 72fa7711d301 / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- vendor_model_list

<a id="canonical-13dc07330e648eca7218d58d42ac422f4f1ab7262910fbf41b1eb3d553947a2c"></a>

Type: `"list"`. Computed.

List of supported hardware vendor and model for this certified hardware.

<a id="canonical-db5f22aea98c95ce2bc8059b30976a35e1f75a9b5320952a048b368bb535c42d"></a>

## Direct properties — vendor_model_list / 72fa7711d301 / 3

<a id="canonical-bc2f307329047f45a9f35fdae12b94898a30db5c0dc02d39c066ac34d608b449"></a>

<a id="canonical-c75c5489da2b8d995190ec66e9e03d83533c4d243e766621ab45d9ad35eee55f"></a>

## model property — vendor_model_list / 72fa7711d301 / 4

Type: `"string"`. Computed.

Hw Model or instance type from cloud provider like number of interfaces, vCPUs, memory.

<a id="canonical-318277bd5075e0a985c701d7dc6896b83b715b9fafb5f50c050c0f89ad4a66c5"></a>

<a id="canonical-25641900fa50ad920c4aa44362c237facb8035a6e2588ac6b8a0cd06b15a3bd3"></a>

## vendor property — vendor_model_list / 72fa7711d301 / 5

Type: `"string"`. Computed.

Vendor could be F5 Distributed Cloud, Dell, Cloud provider like AWS or Azure.

<a id="canonical-d7636d3ffc54f85c541ccc81f24c0e045b200c916455b3613f2473106adac320"></a>

## Next pages — vendor_model_list / 72fa7711d301 / 6

- [Property reference](data-sources--certified_hardware--reference--group-001.md#canonical-3272d898a352fe1b58f8ad58c53ac9f593ec14e32669bfca47c3dd9511d464e1)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
