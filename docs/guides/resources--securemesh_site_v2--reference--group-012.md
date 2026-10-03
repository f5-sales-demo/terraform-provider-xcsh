---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-0010331010102203-0222122220211001-2120002321220023-2202322230333321-0131321221001110-2310321121000332-2123323012113022-0323011021312100"></a>

## local_vrf.sli_config.static_routes.static_routes — static_routes / 231003203323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- local_vrf.sli_config.static_routes.static_routes

<a id="canonical-3231310330232213-2320233132030030-3323331013223103-0102002031032010-1113321301311033-1223210301310111-1333101111133212-1330320303310201"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130031123332312-0202231302321312-0320231321301011-1021121310031101-1230211310131121-3311131210130311-1132032202221200-1122321200010211"></a>

## Direct properties — static_routes / 231003203323 / 3

<a id="canonical-0300233110330131-0321230130313120-1120002302213021-2130121020221211-3032231132330000-1331302201322032-1310111310231120-2201323213212133"></a>

<a id="canonical-3013210102202032-0100133310321012-0103001310213110-2330033333312312-1030322002112110-3121302013132020-3131232110033331-3123321120213211"></a>

## attrs property — static_routes / 231003203323 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-1222221301321120-1200310123133111-0213323301322130-2003112113202102-0121012030201313-2310012320203130-1031203233313201-0131332200013123): complete subsection reference.

<a id="canonical-0103002332302122-0213013322233333-3301203212203330-2201221101021130-0310302101003202-0310232020033310-1110100302223320-3332112313110121"></a>

<a id="canonical-0232123022112133-2100231312200322-3012102300103122-0303333231323010-1112202133003213-2031202001203333-2200020133321132-1212033133310122"></a>

## ip_address property — static_routes / 231003203323 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0131311032003220-2122201031322012-1322120301133133-1031001322103231-0131003021130231-2201102310120110-1123032223011220-3021020303130301"></a>

<a id="canonical-0300112320100321-0013122332123002-0130213133232011-0031130001012131-3022032301020110-0302321010201332-0012233200023222-2022322200220101"></a>

## ip_prefixes property — static_routes / 231003203323 / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011): complete subsection reference.

<a id="canonical-1032022111211323-2100103100000010-0023203120212002-1203133313331133-2202220221323012-2223022021102220-3233332300332311-1321312223332332"></a>

## Next pages — static_routes / 231003203323 / 7

- [local_vrf.sli_config.static_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-1222221301321120-1200310123133111-0213323301322130-2003112113202102-0121012030201313-2310012320203130-1031203233313201-0131332200013123)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1222221301321120-1200310123133111-0213323301322130-2003112113202102-0121012030201313-2310012320203130-1031203233313201-0131332200013123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303331231320203-1021211031211213-3210010001320103-1122331130211311-2020012201010222-0132233302333012-3332101210300103-3031303113123201"></a>

## local_vrf.sli_config.static_routes.static_routes.default_gateway — default_gateway / 123012021302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- local_vrf.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-0030001023012333-2123303232313020-2301301122300030-3133301313031201-3301302310032331-1302003322131000-3323112230211322-3030200201030121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

Upstream description:

This can be used for messages where no values are needed.

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
default_gateway = {}
```

<a id="canonical-0211300012132003-1232022131113222-3102012200103111-3132200232230111-2201131223310321-3031031102123212-3302011213033032-1102312103121122"></a>

## Direct properties — default_gateway / 123012021302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203330213000200-3001031303020310-0113103332123133-1203211332111001-1302322302110221-2220221000022132-2230131222202201-0120333011132233"></a>

## Next pages — default_gateway / 123012021302 / 4

- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021220101021031-3031311120310231-3031321210211101-2032033133310212-1203200011332232-0332132303113032-3133101102331202-1121311133120030"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface — node_interface / 201110223100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- local_vrf.sli_config.static_routes.static_routes.node_interface

<a id="canonical-2232033031322202-3212121121011002-1120021113113032-3030322223131103-2131113312202110-3331100022300302-2303122303213312-1210203222322300"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330101331301213-3100123232303131-3322112223333333-1330012130231121-3323111030300213-1220333312213130-1322123011030220-2001123232112311"></a>

## Direct properties — node_interface / 201110223100 / 3

- [list](resources--securemesh_site_v2--reference--group-012.md#canonical-0102023020212301-3210010200212333-0320310313011103-2102112123100331-2301120110233112-3210011311112122-0032033303133123-0303103323221123): complete subsection reference.

<a id="canonical-2221322121132223-3130003103313033-2023321220231233-0131103333320012-0203232323022202-2031320032031200-3020022220221310-0132110211023012"></a>

## Next pages — node_interface / 201110223100 / 4

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-0102023020212301-3210010200212333-0320310313011103-2102112123100331-2301120110233112-3210011311112122-0032033303133123-0303103323221123)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0102023020212301-3210010200212333-0320310313011103-2102112123100331-2301120110233112-3210011311112122-0032033303133123-0303103323221123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331012112313101-1000223323201310-2130030300121231-3121121102112021-0013311313330121-1110202233002120-3320002110310022-0213122333002013"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface.list — list / 103033110221 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-3033331131102212-3010312310202210-0101100012010111-0321331213002012-0123002213302021-0212112323000123-3020032220212330-0301012310322212"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301133212003003-3130233031200213-1002131310131213-0012131320113003-2210232123133122-1021013030121001-3011110022221313-3133000303113121"></a>

## Direct properties — list / 103033110221 / 3

- [interface](resources--securemesh_site_v2--reference--group-012.md#canonical-1130110231120131-0100113210122102-1003331032013330-2101001331221111-1121330010202101-3300032213031131-2332111130100122-0132201331033103): complete subsection reference.

<a id="canonical-0311031232103310-2201103130322230-0313022333023320-3212130330331222-0301002032101101-3130123233030330-0203113320321031-1222201100300311"></a>

<a id="canonical-1121230022322000-1103102122131233-0021200113113120-3330103230113013-1313221323302320-3103012203010201-2210030100230020-0321300211300323"></a>

## node property — list / 103033110221 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2132202110231130-1121130232330202-1002302300101230-2312012302223300-1023133123032332-0321110210202012-3102323311320100-3303000022130321"></a>

## Next pages — list / 103033110221 / 5

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-012.md#canonical-1130110231120131-0100113210122102-1003331032013330-2101001331221111-1121330010202101-3300032213031131-2332111130100122-0132201331033103)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1130110231120131-0100113210122102-1003331032013330-2101001331221111-1121330010202101-3300032213031131-2332111130100122-0132201331033103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112300122321113-2033012330120310-3232130312120212-2221003203310233-2132232030122321-0021033113221233-0030103000231223-1102121122003310"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface — interface / 120103230310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0323230100211331-3201010220120131-2223322013313002-2002001121303002-0303101233033103-2212023022321213-3011121132202210-0210310233011213)
- [local_vrf.sli_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-3112302003210321-2322123113322011-1033133112020001-2231333223231031-2303322230311333-1210203323200112-1020213323000221-2133021211123033)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3231230102310001-2103103230203001-1133001122200301-2202110123313112-3200000232021122-2101203233022021-3021300020220330-0232130332000011)
- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-0102023020212301-3210010200212333-0320310313011103-2102112123100331-2301120110233112-3210011311112122-0032033303133123-0303103323221123)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3222221200101033-1211120322310023-0322303333120310-0102010033113013-0020023021023102-1030010301030223-1231003322330222-1132102300300232"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202122221232320-1320013110130020-2211120113212102-2000120012330032-1132232330230210-2320223120323021-0021113000023121-3033011230210112"></a>

## Direct properties — interface / 120103230310 / 3

<a id="canonical-2313312210021202-2033001221200233-2130231231113030-3033310112223122-3121121220013102-0133120121112103-3012311031103030-1221012003000020"></a>

<a id="canonical-2303120111132031-3310322311100310-0012221101313102-0023210111020320-0201312013331333-0321222133311023-0301331133311013-3230313100202212"></a>

## kind property — interface / 120103230310 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2102232000121202-3002112320011330-3313010033200112-1102111333130212-0132210310123213-2232323201001131-2011103213003031-0323013233031123"></a>

<a id="canonical-0220101230302011-1200222210302123-1102000231221222-0110333003233001-0323010122303322-0101303102202233-3213001222302120-2231330203201110"></a>

## name property — interface / 120103230310 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0020102230031230-3130231001232321-0323023201111230-3113211133011021-1333011302122123-0021003320312203-2210331102011131-3303210102333333"></a>

<a id="canonical-2300312032023133-2213300102201111-3201231203111032-2200231202330101-0333112311333233-0301220232033231-0311300131010200-1231130000330332"></a>

## namespace property — interface / 120103230310 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1302031311010221-1301201132031123-3212120321013312-1133232212113211-2323313011302321-2231131103133313-2130311222320122-2030011313011123"></a>

<a id="canonical-0300333312213332-2331312000032012-3123302120303223-3202322023300321-0332102222122231-0133102221032021-0101002303130312-1200102032131020"></a>

## tenant property — interface / 120103230310 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2223312331232320-3312100010231130-1100223222033022-2110000212211230-3212232120001000-0323213210113331-2200023110121013-2100003313221132"></a>

<a id="canonical-3130102333321021-2212130221222232-2303332302111103-1011202222012213-0211102332111232-1000113010231310-0302230123311113-2010001121112033"></a>

## uid property — interface / 120103230310 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2013212001223311-3033212111103012-3023312002031320-2012320031031022-0200103233100132-3113131103230002-3103022111112330-3311222211110030"></a>

## Next pages — interface / 120103230310 / 9

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-0102023020212301-3210010200212333-0320310313011103-2102112123100331-2301120110233112-3210011311112122-0032033303133123-0303103323221123)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332002202121300-1111213331310223-2300210003001031-2311222133002130-1030201113300231-0013011022303001-3203021031111030-2113000231333101"></a>

## local_vrf.sli_config.static_v6_routes — static_v6_routes / 103111301033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- local_vrf.sli_config.static_v6_routes

<a id="canonical-0303001002330011-2330332120220020-3223120101123222-1112201302002220-0310101211033123-3231130332101112-1213221111012000-2211230313202133"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203332312331132-0211130233100200-3103213032302211-1202312200213213-3022333230100122-0000212123322223-1031111221122100-2303301313313121"></a>

## Direct properties — static_v6_routes / 103111301033 / 3

- [static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212): complete subsection reference.

<a id="canonical-2111323212132230-2112302023001200-1122121313101321-2033022031022232-1222120020230210-3332001030300012-1233100202002200-3110332030330023"></a>

## Next pages — static_v6_routes / 103111301033 / 4

- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222003201311131-3112221133231311-3321312200312131-3121201331133313-0311012010210132-0201010030111213-1022321120001001-2133101203023011"></a>

## local_vrf.sli_config.static_v6_routes.static_routes — static_routes / 220221123330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- local_vrf.sli_config.static_v6_routes.static_routes

<a id="canonical-0321111330330212-3202110332033300-1310333010022021-2232031012230012-1211121312002223-0031031303133213-3230322120203133-0200312101303322"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301200121210112-0300200332100021-1103122012000033-0301300000312222-3011121312021331-0301212330311123-0123312322222033-3123111213120102"></a>

## Direct properties — static_routes / 220221123330 / 3

<a id="canonical-1023032310320231-1203230302022233-2033113020032213-2331133122231301-1120301110321120-0211311323030013-2230331300322121-3102212333211232"></a>

<a id="canonical-1233313333223123-2302010112110130-0033020201120312-3121020233303022-1100321013013032-2323032130231333-3301220333330212-2210033210113101"></a>

## attrs property — static_routes / 220221123330 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-1220231010001020-2320320123020211-2133332032032202-1331213101223030-1102302301130123-2112112111310002-0020030003100311-1213302303300120): complete subsection reference.

<a id="canonical-3012323111223321-0030103313021223-1120100012302322-3023310233110031-0133103023030021-0332020300303330-1000001013213213-3132331111013021"></a>

<a id="canonical-1110013223311233-1202021003330332-1211230333310031-1301231003201331-2231233333121103-2031012033311201-0113202023032101-0100030210100303"></a>

## ip_address property — static_routes / 220221123330 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0033233210133320-1310133030202001-0220031110323000-3112030121223310-0213010001132121-0020223022001012-3010112333202101-0323130020023122"></a>

<a id="canonical-0133130133010231-1132000111220200-0132211200302211-0131300322202002-3321330320022120-2332030323121303-2122020330033330-3213201313102232"></a>

## ip_prefixes property — static_routes / 220221123330 / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023): complete subsection reference.

<a id="canonical-1011323020201102-3111001300220330-2101132121112020-3033111230232133-0022022233013002-2220313203331111-1330332332030131-2030032302130220"></a>

## Next pages — static_routes / 220221123330 / 7

- [local_vrf.sli_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-1220231010001020-2320320123020211-2133332032032202-1331213101223030-1102302301130123-2112112111310002-0020030003100311-1213302303300120)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1220231010001020-2320320123020211-2133332032032202-1331213101223030-1102302301130123-2112112111310002-0020030003100311-1213302303300120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331312321122101-2102203311132011-3022303022200300-2320322311112011-2222330210321112-3103232321133000-1203132210303122-2001200221102230"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.default_gateway — default_gateway / 302010233231 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- local_vrf.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-3111012110000222-3103332312111120-0203220332331203-3001111321103310-2302233332131012-1221111221033021-1032322010113033-2013203133233203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

Upstream description:

This can be used for messages where no values are needed.

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
default_gateway = {}
```

<a id="canonical-3303203110013320-0001211323221303-3223323333122122-1103320313113011-1211002320133320-1302302132122021-2002303121123233-3032230310213113"></a>

## Direct properties — default_gateway / 302010233231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011331011311322-2030002333123003-0332123202221123-0311313101002033-1033300131103320-3323012101110201-3203202330012010-1113210301032302"></a>

## Next pages — default_gateway / 302010233231 / 4

- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232320330211110-1301012200332222-2023210213121323-3110020232132221-0113211101102203-1313033000002331-3011121220212310-3300333203320233"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface — node_interface / 031012131322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-1120331223212220-1002022123222022-3113111030110021-2232203003001011-3210331103103200-1031221301132003-1031110020213202-0122313322202012"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133101101002323-3330011130212013-2310120223021222-2112100313123331-0020030012131113-1121202231313121-0210013233121301-0323223233000312"></a>

## Direct properties — node_interface / 031012131322 / 3

- [list](resources--securemesh_site_v2--reference--group-012.md#canonical-1033331322002333-2131131330201323-3022032013021033-2213223120110031-2113033022203311-2030112213233222-0100311313122301-3221232120103301): complete subsection reference.

<a id="canonical-3330320111122110-1223310022000021-1222300202121331-0331312012220002-1012122100101302-2213011233233122-0311120131220120-0230123310012322"></a>

## Next pages — node_interface / 031012131322 / 4

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-1033331322002333-2131131330201323-3022032013021033-2213223120110031-2113033022203311-2030112213233222-0100311313122301-3221232120103301)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1033331322002333-2131131330201323-3022032013021033-2213223120110031-2113033022203311-2030112213233222-0100311313122301-3221232120103301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002003323213102-1123220223312130-2223003303010101-0010000122123123-0303020330123200-1103301000210311-2033310003130330-1002322001301031"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list — list / 011022011222 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-3031111100323322-1003211101031210-3202121122321110-1332213310232301-0203111032220010-2320231330332120-2311322330101311-2132332230110120"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211210123331020-1133001032031300-2132223032123123-1201301310102202-3022313310012213-1000013021211100-1313011112200130-2101230020013301"></a>

## Direct properties — list / 011022011222 / 3

- [interface](resources--securemesh_site_v2--reference--group-012.md#canonical-1012310210121023-2123121011130310-1033321320002210-3023333323013011-0231323021112311-3013021312101210-1100203112003330-0211201221110132): complete subsection reference.

<a id="canonical-3101123220130302-0012223211321202-2212322331013230-1213320012003313-1113101103121113-1203303103331310-3020011202003133-2202303321000321"></a>

<a id="canonical-2303311023122021-0103023130023333-1103031323230202-3312332011203231-0021021030121300-1332022201200103-2202303301200133-3130300203120121"></a>

## node property — list / 011022011222 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311102332231013-2130231320020030-1322113223122011-1031201312112103-1233311110213212-0213322113120113-3103220010210022-3012031001221222"></a>

## Next pages — list / 011022011222 / 5

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-012.md#canonical-1012310210121023-2123121011130310-1033321320002210-3023333323013011-0231323021112311-3013021312101210-1100203112003330-0211201221110132)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1012310210121023-2123121011130310-1033321320002210-3023333323013011-0231323021112311-3013021312101210-1100203112003330-0211201221110132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223003210133330-1120200311301030-3232333033312200-1330223131131133-1013001223332001-2330121131331122-1311323123301112-2130111301132111"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 022002230101 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.sli_config](resources--securemesh_site_v2--reference--group-011.md#canonical-1003033010321202-0033020021201100-0330201010332332-1101302032222123-3100130130311320-3331333321202210-0320333303212021-1033121020003000)
- [local_vrf.sli_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3001132010301012-0220231311323020-0311322102311121-0121033001331230-2122211132122333-0102332113200201-2033302021232330-0312312003001000)
- [local_vrf.sli_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2110222100132232-2101010123323212-3313331003332003-0312321330102112-2312013033233330-0231211020033110-1031002210113120-0313110330310212)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3021030203330201-2202101133033231-2020210222100020-0331231202100132-0221223023102102-3022301003200310-1331322321021120-3131220111312023)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-1033331322002333-2131131330201323-3022032013021033-2213223120110031-2113033022203311-2030112213233222-0100311313122301-3221232120103301)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-2110203031100002-1301103321000003-2013113011122303-3101103322123012-2321222221201100-1223330212203312-0333103313123032-1033211032310031"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211112002110303-1113111031333331-0031001312011120-0201033333203212-0203223233101233-1121002301311320-3121320322212212-0231111222230030"></a>

## Direct properties — interface / 022002230101 / 3

<a id="canonical-3023212112331000-3332322122021213-0213120231133022-3313030312323223-1331021012120310-2012201200312103-2223030000310302-2221032000300021"></a>

<a id="canonical-0310302331020310-0331301203333332-1102001002321123-2211231212102320-2332100213033121-0122313020323113-3010011201130021-0011131001032202"></a>

## kind property — interface / 022002230101 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3113212012130100-1112203121030021-3111231232303132-3013120321120333-1123101130002001-2303022311323320-3103111110212303-3200210122103310"></a>

<a id="canonical-0221010111232222-1313233300121133-1133201101100003-2000313033321031-0033303031100332-1323000210023330-2121100232010023-1030301010102002"></a>

## name property — interface / 022002230101 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3130222220013332-0200211233020103-1131303320202001-0320303331032211-0001000123303221-2332331323301131-3311333301332310-0233120121032310"></a>

<a id="canonical-2111303312033210-3332201023210021-2221113202031323-2110112332211212-0332301322120131-3031332201022102-1011032033113220-3111101321300120"></a>

## namespace property — interface / 022002230101 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3001313212231110-0020231100332213-1221300202320112-3323131031231202-0300211103112221-0222211103100113-2133223323330020-3223311100033210"></a>

<a id="canonical-0102320032130121-1302001231021003-2322203130031022-1223210203111313-2100023112103222-0312302201220332-0330331223222032-3123323223231303"></a>

## tenant property — interface / 022002230101 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223130033130110-3023021012330030-3212310112222033-3101211223323223-1223120330212201-3221210203222110-3331320110130333-0032121112131130"></a>

<a id="canonical-1323201022001003-0030012100233021-0133021102233230-0210031200312032-3233030220030031-3011230301211102-2101010101222022-2331003102222112"></a>

## uid property — interface / 022002230101 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0103312333123203-0203212103103312-0233102211121320-1003223332330133-2322301102113123-3322133213111011-3022011110021030-1130300011020230"></a>

## Next pages — interface / 022002230101 / 9

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-1033331322002333-2131131330201323-3022032013021033-2213223120110031-2113033022203311-2030112213233222-0100311313122301-3221232120103301)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131201102301310-1032231002121213-0013011332132200-0331000301100223-3222332030200310-1203221323033331-3012323123103221-3130030300033302"></a>

## local_vrf.slo_config — slo_config / 331030003100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- local_vrf.slo_config

<a id="canonical-0023201332223302-0123003312310202-1011110332203000-2033320102323201-1012322100101103-1222331131223002-3000230221222010-1111322013103330"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
slo_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130023121211230-3312120320330300-2210301221310032-0200110322311103-3013102332032120-3111113302131202-3313011122303023-3212001002102030"></a>

## Direct properties — slo_config / 331030003100 / 3

<a id="canonical-1233112121331011-3223100010222210-2011210232333203-3330013230322302-2303111220011102-0210211233222012-1102101332002132-3311102003312021"></a>

<a id="canonical-2130113121011012-0332301320031221-3320212232300302-1112032333311331-3231030321202030-1132121103100102-0201300331232013-3023221030023232"></a>

## labels property — slo_config / 331030003100 / 4

Type: `["map", "string"]`. Optional.

Add Labels for this network, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0120310020200323-0122333221010312-2221213121211001-0233013331302033-2301210121011113-2301110230123102-2011201222320022-0100300233113311"></a>

<a id="canonical-3031232102113112-3012030032303110-3113232311311203-0303121312100000-0102211302002122-3331132310131121-1113001003200322-2000321311333123"></a>

## nameserver property — slo_config / 331030003100 / 5

Type: `"string"`. Optional.

Optional IPv4 DNS server to be used for name resolution.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [no_static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3010131300322320-1003302333332300-1020211101201223-1330230321111000-3302100331030033-2131211010320101-0113010111321222-3110323122220301): complete subsection reference.

- [no_v6_static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0012230222212130-0202203112130311-0010030110131100-0230201032312322-0332203222000101-2001310101203203-2120032012002122-3001033010123112): complete subsection reference.

<a id="canonical-2312230010322012-0111232111233023-3321003321230110-2102313323320322-2012313033020020-0311310322221222-0101230110331313-0311100122112202"></a>

<a id="canonical-3000100321121310-3331120030303130-1003330233103330-3231310322311310-2123011113002100-3230300230233032-0211233333031111-3000002310013213"></a>

## secondary_nameserver property — slo_config / 331030003100 / 6

Type: `"string"`. Optional.

Optional Secondary IPv4 DNS server to be used for name resolution.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210): complete subsection reference.

- [static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020): complete subsection reference.

<a id="canonical-3203210313231201-1032100131022131-0312202012012021-1010233122333200-3103101131210112-0023302230020231-3221101220202003-2220210130003312"></a>

<a id="canonical-1003330300002313-1333202221330320-2321000313302330-1223312311213321-2312313123022121-3002002211102302-1203102032033213-3212102213322123"></a>

## vip property — slo_config / 331030003100 / 7

Type: `"string"`. Optional.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3121311333002220-2213010333311031-0011120321000211-3330103222323021-0121220122221120-3202022131230100-1112113303310302-0130111030232310"></a>

## Next pages — slo_config / 331030003100 / 8

- [local_vrf.slo_config.no_static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3010131300322320-1003302333332300-1020211101201223-1330230321111000-3302100331030033-2131211010320101-0113010111321222-3110323122220301)
- [local_vrf.slo_config.no_v6_static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0012230222212130-0202203112130311-0010030110131100-0230201032312322-0332203222000101-2001310101203203-2120032012002122-3001033010123112)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3010131300322320-1003302333332300-1020211101201223-1330230321111000-3302100331030033-2131211010320101-0113010111321222-3110323122220301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310100010111213-3000203310110121-0313321102000321-2301132311123312-3303031332330231-1130321130010220-2202313322033222-1231003232022131"></a>

## local_vrf.slo_config.no_static_routes — no_static_routes / 200313210303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- local_vrf.slo_config.no_static_routes

<a id="canonical-3232121310133000-2320211000131013-2103112100330113-3010310300301101-3112112032121321-2233010112302102-3233300213031313-1011120332203030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no static routes.

Upstream description:

This can be used for messages where no values are needed.

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
no_static_routes = {}
```

<a id="canonical-0231333302313200-0032221121023202-0233031000122013-2320212222303201-0213313130333103-2131122210331013-0000132010123101-0021112110203020"></a>

## Direct properties — no_static_routes / 200313210303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010200103221030-2100020103311033-2321021111103002-0212131133120222-0302311202113320-1330311213032220-0310233203113223-3322231020312030"></a>

## Next pages — no_static_routes / 200313210303 / 4

- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0012230222212130-0202203112130311-0010030110131100-0230201032312322-0332203222000101-2001310101203203-2120032012002122-3001033010123112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031222002122000-3131310131233102-2230132232333303-1203310331023121-1212202113202101-2230302200103312-1233101323111132-2222322230103233"></a>

## local_vrf.slo_config.no_v6_static_routes — no_v6_static_routes / 321203313121 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- local_vrf.slo_config.no_v6_static_routes

<a id="canonical-2102330223033123-2221322213230122-3003302300012111-3032203132221100-1222110300020133-0300013211230310-2322223322333232-1200101021331033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no v6 static routes.

Upstream description:

This can be used for messages where no values are needed.

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
no_v6_static_routes = {}
```

<a id="canonical-2212103021323320-0011211112113131-3213200102202031-1232232111003033-1110101330310302-2120203020132230-0131332320012031-1120000020133012"></a>

## Direct properties — no_v6_static_routes / 321203313121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000103112031221-0102100322103222-2220032030121120-0120302122120120-3020313223310231-0002213203130010-1023013212101101-3100112230102012"></a>

## Next pages — no_v6_static_routes / 321203313121 / 4

- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033320021122132-0301022221033300-2110330030131320-2310203312232213-3130233313012300-3111023220213310-0322011021100332-1311011000003123"></a>

## local_vrf.slo_config.static_routes — static_routes / 100322300102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- local_vrf.slo_config.static_routes

<a id="canonical-1321212023130110-0212000110101322-0310101102113012-0202301133200112-3313033200122013-0202323232333032-2231310201122331-1213031300030222"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233222211330122-0222302333330212-0000133122333310-2221122231301010-2011123030203220-1031133200002203-3123333211012100-1123133133021323"></a>

## Direct properties — static_routes / 100322300102 / 3

- [static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200): complete subsection reference.

<a id="canonical-2131203213101001-0321023110113012-1013131102120322-2322230321323320-0111233133030003-0113021030033131-0313120000322133-0310303233121201"></a>

## Next pages — static_routes / 100322300102 / 4

- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001200010221303-2330023313133202-1102300013312001-1032123313131233-1011322202022313-3003120112003033-2223122033330101-1133231002230102"></a>

## local_vrf.slo_config.static_routes.static_routes — static_routes / 003303013121 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- local_vrf.slo_config.static_routes.static_routes

<a id="canonical-1032020232102223-3221010302110022-2031021002310332-0000002012222123-2102233103231101-1100110212200002-2323103121102222-2111102200020112"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310012302023122-3231233103132212-1110110331221130-1212302202013011-1021303130001113-2101011103221212-2002233021221310-3033311002011313"></a>

## Direct properties — static_routes / 003303013121 / 3

<a id="canonical-0201022233020313-1102200132030322-2011003131301301-3101330132012001-2211120131323230-0330113330123120-0022020130233030-0103231121003132"></a>

<a id="canonical-1003113220211103-1031202022301131-0000132232330120-1023002101103133-3110030102021030-1312311301220333-1221331032221322-3123002121220223"></a>

## attrs property — static_routes / 003303013121 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-1010313030211131-2011012032320033-2033001112122133-3112003321210322-1112323103311221-0313222030300213-1232231011011031-1132032100211323): complete subsection reference.

<a id="canonical-0201323221111302-0220033201002320-1102102031001000-3002203001330230-3321200213122112-2203323121020013-1220123233010320-2020310031202302"></a>

<a id="canonical-3003321010113003-0122202111022022-3323112222001211-2321202012203110-2133220303223023-2001220003123131-2102320122323032-0312001013331232"></a>

## ip_address property — static_routes / 003303013121 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3122213110000311-3100021030322203-0333331200111233-3002300202032223-2112320330233211-3002033232230212-3302002222103133-1023011012102221"></a>

<a id="canonical-0110332200113110-0120311311122000-3221021220122322-0332303132323010-3001020220310303-3001112122010013-3313032111131331-1033130321001013"></a>

## ip_prefixes property — static_routes / 003303013121 / 6

Type: `["list", "string"]`. Optional.

List of route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122): complete subsection reference.

<a id="canonical-1230333010302203-0200213302023311-0131322023311022-2132003333331203-2012311322312330-3131013003310223-3201323300333201-2203002033233033"></a>

## Next pages — static_routes / 003303013121 / 7

- [local_vrf.slo_config.static_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-1010313030211131-2011012032320033-2033001112122133-3112003321210322-1112323103311221-0313222030300213-1232231011011031-1132032100211323)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1010313030211131-2011012032320033-2033001112122133-3112003321210322-1112323103311221-0313222030300213-1232231011011031-1132032100211323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001001231123110-1210110030333312-0311122230203230-3221021211131121-2222323230222232-0230110231033210-1313101232200120-1220312212313230"></a>

## local_vrf.slo_config.static_routes.static_routes.default_gateway — default_gateway / 130212302202 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- local_vrf.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-2103312212110233-1203202030300230-3003200013331223-3030203012212000-1303023023330100-3323113012033021-0002213211332221-1033333220133331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

Upstream description:

This can be used for messages where no values are needed.

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
default_gateway = {}
```

<a id="canonical-1223132010333113-3300012312222133-1213100321311121-3013232322102332-3130133113213000-1321211020313323-3010231122022103-2020312331131110"></a>

## Direct properties — default_gateway / 130212302202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223313231320011-1022122023102230-0122310132103232-1223331002212211-3032331331222010-2111201003020021-0312300103323222-3232122333232330"></a>

## Next pages — default_gateway / 130212302202 / 4

- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220032033303232-2120301320102322-3313230211230021-1313003103332310-2030302110133330-3012311001223132-1330330300001211-0020121133202200"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface — node_interface / 003102313110 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- local_vrf.slo_config.static_routes.static_routes.node_interface

<a id="canonical-0133333012103113-0321122203233210-3011010222020313-0120321220120212-1110331221213210-0203023222322132-1211033123321023-0021033312232103"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001311220221210-2100332000233032-1220333111122231-2210113231011203-3022233310220021-3231320310200210-2123202233223001-1003322233222033"></a>

## Direct properties — node_interface / 003102313110 / 3

- [list](resources--securemesh_site_v2--reference--group-012.md#canonical-0020231101330112-3322130101230311-1033000012133232-2011131313233010-2013002233110303-1130331312330132-1022322332221001-1332033111030113): complete subsection reference.

<a id="canonical-1233120232223332-2202320110000123-2303120102312232-3321130232030312-3323232121031332-0122333330310000-2312211012010202-2221000230213213"></a>

## Next pages — node_interface / 003102313110 / 4

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-0020231101330112-3322130101230311-1033000012133232-2011131313233010-2013002233110303-1130331312330132-1022322332221001-1332033111030113)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0020231101330112-3322130101230311-1033000012133232-2011131313233010-2013002233110303-1130331312330132-1022322332221001-1332033111030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221121303031300-3313302221123110-1232002233023311-2033322303123120-2000131021331112-3030310310211211-3313232321322110-3003102223011033"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface.list — list / 000120011132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-0211311020320032-1012301333133023-1122231122203202-0300332331012213-3200131120132331-0213001023233000-0212231223131332-3130100303033023"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030333100030003-3002313203201303-1233321233211011-3212012113030102-0102211203003102-0323112210011121-3103230233311111-3331120031332020"></a>

## Direct properties — list / 000120011132 / 3

- [interface](resources--securemesh_site_v2--reference--group-012.md#canonical-2302111313203322-0232000210330122-3013320321311122-0131330222331230-3202313302113233-3313213023223332-1102200201110123-2211200203333332): complete subsection reference.

<a id="canonical-3021320000312222-0001103332330300-0222203210210033-3333000220022132-1131233021302120-2321132102201203-1103111230301113-3023300321220200"></a>

<a id="canonical-3323200322123320-1131221303221332-2031320323100331-3033113311200020-0101303112131320-2332213100003233-0211200123333122-3013113331020133"></a>

## node property — list / 000120011132 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2300121103100030-2323210110032211-0132202312003302-1020010120320033-2321221303312132-3130213123002110-3330022001222012-2121220022032313"></a>

## Next pages — list / 000120011132 / 5

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-012.md#canonical-2302111313203322-0232000210330122-3013320321311122-0131330222331230-3202313302113233-3313213023223332-1102200201110123-2211200203333332)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2302111313203322-0232000210330122-3013320321311122-0131330222331230-3202313302113233-3313213023223332-1102200201110123-2211200203333332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222100002021231-3030200113323211-2102233020002321-3313311333001002-2033322200111203-0321032122123233-0323031122331100-3133131111301033"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface — interface / 102011012211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-1323000113011001-1312003312320210-0110232003230220-2111012303211122-3102232201300223-3203203110321331-3131122032110013-0112120321002200)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3211220231201301-0100330212203123-0023001200223110-2122301010000023-2120010301012031-0300203221223221-3021132110013321-3113011322100122)
- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-0020231101330112-3322130101230311-1033000012133232-2011131313233010-2013002233110303-1130331312330132-1022322332221001-1332033111030113)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-0310231231001113-3201101320022232-3213101202000100-0011020220103002-3323000020320020-0231303020133001-3200111323001103-3002230332321322"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021222132101233-3113213102031311-0110012012033033-0210101021112323-3003202112223223-1230003002112031-3022020210122002-3302233211231323"></a>

## Direct properties — interface / 102011012211 / 3

<a id="canonical-1203000202031313-1232203120020023-1100233203110020-0223230011021332-2113202011310132-3322313100133100-1311312003220123-1133020232311120"></a>

<a id="canonical-3132310101101003-2302203002031301-3002302122030133-2113212032113013-1222113210002030-3132222303012320-3111031222303022-1203321113213303"></a>

## kind property — interface / 102011012211 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2020312332211033-2203202021230131-1120220232322332-3313223201301113-3302300130132223-1323022320123312-1300232211200332-3101100111313131"></a>

<a id="canonical-0232030232123302-0112220232331312-2220013203110131-2332320310212223-0311132032113220-0220301323312012-0120223021101133-3132101101233123"></a>

## name property — interface / 102011012211 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0331010300102110-1022210123202130-1021031320033220-2303021213220000-1012222133113111-1230033310202321-3111302322122202-0211300203222012"></a>

<a id="canonical-2120033320032020-3202322212101012-2023003331310301-3022320132320122-3102321113001020-0032131121333132-1232000231112233-0231023322121302"></a>

## namespace property — interface / 102011012211 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2302222202212230-0213013321001303-1300111301113213-0023002031303212-0322311210230130-0200320203223110-3233223103103030-2113310313000213"></a>

<a id="canonical-3102002023231313-3031222131303330-0023321010031311-3223223002220213-0101330111203021-1031333131320331-2322230013200303-2003310113103112"></a>

## tenant property — interface / 102011012211 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2232232220121020-1130230021301032-0021101132113020-3110022201332323-1020013331301133-0321030223332111-1123312100232123-3331000303102031"></a>

<a id="canonical-1230010002232013-1310203110303103-1322201033200130-2203013200110113-2202202213232320-1203300230203100-2232003030213313-0131202020130321"></a>

## uid property — interface / 102011012211 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0011230322303030-1203002220323121-3302032221312322-0022012213213233-1002110303110022-1002213010303233-2012021302113221-0133223112230203"></a>

## Next pages — interface / 102011012211 / 9

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-0020231101330112-3322130101230311-1033000012133232-2011131313233010-2013002233110303-1130331312330132-1022322332221001-1332033111030113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203222010102013-3323030302301130-2021010221011103-1133203321233101-1110311331101120-3030333201130000-2100123213311133-2103233103333121"></a>

## local_vrf.slo_config.static_v6_routes — static_v6_routes / 221313221022 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- local_vrf.slo_config.static_v6_routes

<a id="canonical-0132323220130201-2200123312322222-0120201133210122-2333331103000323-3232100022230333-2001323220020303-3211101120310332-1311023200313210"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313002130122103-2010213120122331-2232200103300222-0313011021020220-3023322123323322-0021133312210322-0231303001233110-3012311302010302"></a>

## Direct properties — static_v6_routes / 221313221022 / 3

- [static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202): complete subsection reference.

<a id="canonical-1303301333323313-2121321311001121-2331200300021213-3110301303220101-0222023302312332-0202300320103012-3023011223110321-1303330231021230"></a>

## Next pages — static_v6_routes / 221313221022 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321020332210122-0032111122112000-1000131232202302-2020302231212132-0102103122221303-2012111020211221-1033312123110232-3220111002302100"></a>

## local_vrf.slo_config.static_v6_routes.static_routes — static_routes / 312210102022 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- local_vrf.slo_config.static_v6_routes.static_routes

<a id="canonical-1103112031031313-3011221211032023-1112111203013212-1313313323120210-2210321301022012-0001331101103203-2102111311021223-0130222233231320"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330211301211203-2001101222213220-1122221201231300-2220110132320202-1011203111032210-0303220030020102-0000321330311222-0101301112021322"></a>

## Direct properties — static_routes / 312210102022 / 3

<a id="canonical-2320101101101002-0231223100022210-1011132330322112-1210130120233200-2231212331221211-1310233320213020-0122333221020102-0310031230231010"></a>

<a id="canonical-2002123312032013-1313003011211220-0023130130320313-3330113302003010-2112133210031220-0302030321001312-3213333320131311-2100022310211101"></a>

## attrs property — static_routes / 312210102022 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-3212031133321311-1231003013302211-2112031202020320-3030232001033202-2313212312031333-3100321031123330-0023221313101222-1201201220312101): complete subsection reference.

<a id="canonical-0333212202020102-0223302303120301-1002112031200300-1031232122112300-1222120302103331-3031330100220322-1320331131110100-0022312123332013"></a>

<a id="canonical-3332101103100212-0211332011302210-0332110120030022-1001120230031022-0222021211020030-1212231330213033-2021122120133123-3001011030002012"></a>

## ip_address property — static_routes / 312210102022 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3222023231220230-3310100032000102-0030230110231113-1102103213131031-0222222200332101-1030312313021102-3301313301110330-1320221331332010"></a>

<a id="canonical-3012200030202201-0333101010103231-3221201131022331-3122000133313003-3220120212132302-2211313222323330-0001120033230023-0031212320320002"></a>

## ip_prefixes property — static_routes / 312210102022 / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130): complete subsection reference.

<a id="canonical-1032221031112032-3011310113020313-3121221202200232-0223112022323300-0232000202221332-2323022210131121-1001200001312231-0220120220211322"></a>

## Next pages — static_routes / 312210102022 / 7

- [local_vrf.slo_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-3212031133321311-1231003013302211-2112031202020320-3030232001033202-2313212312031333-3100321031123330-0023221313101222-1201201220312101)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3212031133321311-1231003013302211-2112031202020320-3030232001033202-2313212312031333-3100321031123330-0023221313101222-1201201220312101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010100323231302-3111121013122102-3102001302123202-1312303101032030-1122003121021013-0321330200231301-3220302223312032-0221120111210231"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.default_gateway — default_gateway / 012131103303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- local_vrf.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-1313001231132101-3303211333312331-2203032322300033-3120111122231313-1023303032202310-3233132200032002-0001130331122221-1310330100110332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

Upstream description:

This can be used for messages where no values are needed.

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
default_gateway = {}
```

<a id="canonical-0021233321221122-3133312021133312-0010000100202022-2210213200211303-0303000030122222-0200300302012323-1211210020031110-0313323033201132"></a>

## Direct properties — default_gateway / 012131103303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013301113013211-0312132331302120-2222100311011033-1232012000312130-2300013211113230-2302312211213012-1131103133103003-3121101132120122"></a>

## Next pages — default_gateway / 012131103303 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322221302201030-0003233213100031-1012000133220312-2123011121233332-0100212001123312-0111332313330000-1133200300123232-3211122020031001"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface — node_interface / 332213310210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-2313301111001102-2230112300320213-1011023102220030-3133113021122120-3023322120130303-3012111023232333-1001023230032232-0123001103000312"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200021333222320-3010022002313320-1220222330010230-3122332311033320-3100313202100030-3023320013001111-2102320232302203-1312002223132203"></a>

## Direct properties — node_interface / 332213310210 / 3

- [list](resources--securemesh_site_v2--reference--group-012.md#canonical-1023111113021133-3311101122122022-3112031222130033-0320102321200133-3030031111122200-1022030330102100-3232120132013000-0312023211221330): complete subsection reference.

<a id="canonical-1100010130320132-2220200132120010-1123320011212212-2123012333123312-1032122022332211-1033220100312020-3130010213002313-2011010111201312"></a>

## Next pages — node_interface / 332213310210 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-1023111113021133-3311101122122022-3112031222130033-0320102321200133-3030031111122200-1022030330102100-3232120132013000-0312023211221330)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1023111113021133-3311101122122022-3112031222130033-0320102321200133-3030031111122200-1022030330102100-3232120132013000-0312023211221330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032131131220103-1011230120101230-1323302031330333-2030202203320032-2213300022101230-1322312020123020-3031223032311333-0323111333022033"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list — list / 102112110230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2030302101310233-2103300220122021-3112200122032011-0023110332323102-3202321201223003-3310213200023333-2301103132023132-0122312213321003"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020113001133123-2222203220111032-1201313122033110-0330323211122132-2330003303022223-0130302023203102-2311223113133333-3131310311003321"></a>

## Direct properties — list / 102112110230 / 3

- [interface](resources--securemesh_site_v2--reference--group-012.md#canonical-2122222222000100-3331202322230202-3133030301110313-0302103313201321-0021332232102111-0023322113201033-3220300222010312-1101323202233330): complete subsection reference.

<a id="canonical-2322332230210310-3322100110211320-2311103113101200-2121303022020233-1022132311213120-2130031013001211-1021303033222010-3000322232200020"></a>

<a id="canonical-0233111123133223-2323323113100012-0120130033320332-0233332132000132-3033200330003213-1310113110033131-0130022212223222-0120333012233230"></a>

## node property — list / 102112110230 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022020122002133-3221211311230000-2012201013302223-3330020003022313-1112101121102332-2200310311031013-3023201131033310-2033100032112212"></a>

## Next pages — list / 102112110230 / 5

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-012.md#canonical-2122222222000100-3331202322230202-3133030301110313-0302103313201321-0021332232102111-0023322113201033-3220300222010312-1101323202233330)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2122222222000100-3331202322230202-3133030301110313-0302103313201321-0021332232102111-0023322113201033-3220300222010312-1101323202233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130013120102112-1221200311102011-0020012220023332-3020033131000233-0103310230033031-3201103000133331-0000300022320232-2331313221221011"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 100112231231 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0030322012011132-0000223213023012-0031232202030221-3113303112222223-2200030021032213-2233131213333033-1013103203001303-3130333110112331)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-2120130200002022-1111320122112233-2012113012012100-2303120231330013-1312122331032013-0133221002311112-2031133002133200-3233212233303020)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-3120311120112332-1101122110331023-0012300230001000-1110333013323211-3302201321302231-0132130202011030-2130200103112332-3111030332120202)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-0000011010311220-2032120202003322-1312130122233321-0200121311110000-2300230002110222-0210300122113010-0033123100300301-3230012113203130)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-1023111113021133-3311101122122022-3112031222130033-0320102321200133-3030031111122200-1022030330102100-3232120132013000-0312023211221330)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-0321003312313320-1123311233202322-2301001101032031-1010303212330122-3111310220031130-0211201022101310-1332011222121223-3300301112303332"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010133332012001-3121332101322310-0130210031301131-1312111213010010-1023012101021203-1202213222230332-1001102313230113-1000220333113333"></a>

## Direct properties — interface / 100112231231 / 3

<a id="canonical-1000231220311020-0213220001112223-0323221201110122-2232311102010323-0222330231020332-2331233033113030-1032212130333320-0232002331111023"></a>

<a id="canonical-0212230202322200-0221011231212001-1130102210110023-3103001031323302-3032111111312331-2333212300220013-3130012332103123-1201001210001033"></a>

## kind property — interface / 100112231231 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1132332233123003-2212201132101212-1003102131323322-2120013223330323-0000332111112013-3131200012031002-2002031210010330-3230333130020110"></a>

<a id="canonical-0301021030230322-2211333323200203-0301123130220323-2121300112312022-3233220033113313-1300332032320011-0211011231123010-3102011101033321"></a>

## name property — interface / 100112231231 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0323320010003222-2202331203132013-0120113211230100-3033030212021122-2203210012232021-3322231102210202-0322133222112322-1030321003333222"></a>

<a id="canonical-2223221002131113-0023130303301201-1232013032032330-1101231110013011-3021213211233310-0030023012010311-1223211310030221-3313122111203300"></a>

## namespace property — interface / 100112231231 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0022233200332313-2120302200002012-1200202233031002-0301220233303231-2230103001000200-3101023230311333-3022321103021111-0210201131012112"></a>

<a id="canonical-1013121123113222-0131110212332113-1312001011120202-0020200331321020-1122012132322032-2113010121132223-2211021132021212-3013032000101230"></a>

## tenant property — interface / 100112231231 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3201211231221112-0231321132012000-3003031201210220-1100020200331130-1233120220113123-1312100100311102-0322130300011021-2030231212123300"></a>

<a id="canonical-0033211033133223-3112212120311322-0211333320101111-2232222310203313-2320000332203120-0003012002101113-2021222231231013-1211313130200000"></a>

## uid property — interface / 100112231231 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3033111312011000-0113021011330011-2021233331302201-1013002000020331-0012233301311003-0332112233120303-3230012032021201-1100012030003102"></a>

## Next pages — interface / 100112231231 / 9

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-1023111113021133-3311101122122022-3112031222130033-0320102321200133-3030031111122200-1022030330102100-3232120132013000-0312023211221330)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121310033000313-0022032301310220-2001332130111003-1320222223001333-2303212113011001-2000023200211002-0000332300121033-3323323133230211"></a>

## log_receiver_with_net — log_receiver_with_net / 312001101031 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- log_receiver_with_net

<a id="canonical-1222131012023103-2300202303002133-3112212323102320-1230201032232133-2332123111310012-0333132022200221-2203011330112320-2223302000230020"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver\_with\_net, logs\_streaming\_disabled\] Select log receiver for logs
streaming with network option.

Upstream description:

Select log receiver for logs streaming with network option.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("use_management_network",
    "use_slo_sli")}
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
  "x-ves-oneof-field-network_choice": "[\"use_management_network\",\"use_slo_sli\"]"
}
```

OneOf alternatives in this subsection:

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-1222131012023103-2300202303002133-3112212323102320-1230201032232133-2332123111310012-0333132022200221-2203011330112320-2223302000230020)
- [logs_streaming_disabled](resources--securemesh_site_v2--reference--group-012.md#canonical-0111001211331333-3333100033323212-2110200312112313-2301000232032002-0331312312200211-1233122002101233-2323212221312020-1333023010232331)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver_with_net {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102112201003103-2123310122020102-3200213222302110-1310133001320223-2032300111112133-2233103003032300-1320002113200321-3200011312021333"></a>

## Direct properties — log_receiver_with_net / 312001101031 / 3

- [log_receiver](resources--securemesh_site_v2--reference--group-012.md#canonical-1120012013013123-3022211010330233-3002013101202231-2213130010033121-3321013301122032-1100011113202113-0010303022202231-1003123003220200): complete subsection reference.

- [use_management_network](resources--securemesh_site_v2--reference--group-012.md#canonical-1000230333312100-2312133011203203-2301313000031303-3213013221120231-3211023003322120-2020313121120300-3121213103011210-0023011121111220): complete subsection reference.

- [use_slo_sli](resources--securemesh_site_v2--reference--group-012.md#canonical-0031021002303321-3112130012003112-1011131101211211-0033112221120300-3311120313031021-3210203333312121-1200221213301132-1033211302020230): complete subsection reference.

<a id="canonical-3211201013020113-1331322023133230-3310101021231021-2223010111232013-1310011321313112-1001211020333103-1310130333210101-1101210221013202"></a>

## Next pages — log_receiver_with_net / 312001101031 / 4

- [log_receiver_with_net.log_receiver](resources--securemesh_site_v2--reference--group-012.md#canonical-1120012013013123-3022211010330233-3002013101202231-2213130010033121-3321013301122032-1100011113202113-0010303022202231-1003123003220200)
- [log_receiver_with_net.use_management_network](resources--securemesh_site_v2--reference--group-012.md#canonical-1000230333312100-2312133011203203-2301313000031303-3213013221120231-3211023003322120-2020313121120300-3121213103011210-0023011121111220)
- [log_receiver_with_net.use_slo_sli](resources--securemesh_site_v2--reference--group-012.md#canonical-0031021002303321-3112130012003112-1011131101211211-0033112221120300-3311120313031021-3210203333312121-1200221213301132-1033211302020230)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1120012013013123-3022211010330233-3002013101202231-2213130010033121-3321013301122032-1100011113202113-0010303022202231-1003123003220200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301002330202021-3031103331103310-2211231023223113-1300313022000312-3001022020113203-0030201112213213-1322221133211101-3023230022321332"></a>

## log_receiver_with_net.log_receiver — log_receiver / 032031131332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030)
- log_receiver_with_net.log_receiver

<a id="canonical-3330103200300113-2010313112323020-3011202121021303-2200223220233222-1030223101122202-0022302311103032-3033033011121113-1031111200131121"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323320101113011-3222310122332301-3202231333203221-1231323020133111-3133230001012211-1021200130013330-1332222222202301-2330022310333331"></a>

## Direct properties — log_receiver / 032031131332 / 3

<a id="canonical-3231032321122013-2000303131120300-2122300022211023-3022333131203011-0203212231302233-1001321210023302-3122111012103003-2323113110013210"></a>

<a id="canonical-0023310210111110-3131330032233120-0023132312101232-2201100001301031-1233302213101300-0030013001132330-3303103013001212-1013223101102212"></a>

## name property — log_receiver / 032031131332 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1213120300332313-2330131031330231-0121021231311001-2321011201010121-2330033130100222-3231300201032031-0031230330232301-2211002321321103"></a>

<a id="canonical-3012313323212313-0231102221210011-0011231232320200-2220220120301022-0300331323223332-3033303333313003-2203132132000103-2112213103333032"></a>

## namespace property — log_receiver / 032031131332 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3012310101132323-0323331133012022-3323330232210103-0123030230331022-1332011300213320-2110310112013210-3122201302112023-3111311200232110"></a>

<a id="canonical-1333121002312133-0132312000211211-1312030120000012-0213121001301100-1311003221022020-1210120321022332-2123132320102023-1222221012013332"></a>

## tenant property — log_receiver / 032031131332 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1233001020310301-2131302122011122-0133231102010102-3110222312131232-1210030111010110-2321110131320021-2113103111202220-0033310332313010"></a>

## Next pages — log_receiver / 032031131332 / 7

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1000230333312100-2312133011203203-2301313000031303-3213013221120231-3211023003322120-2020313121120300-3121213103011210-0023011121111220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231130303001111-2311123332133200-0122031210222320-2212113221130233-3302102203023013-1111211022111300-0301230012133112-2230131130012200"></a>

## log_receiver_with_net.use_management_network — use_management_network / 302312021211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030)
- log_receiver_with_net.use_management_network

<a id="canonical-2033320112213113-3132213100330323-2202012010033003-0322311102120110-3023220123020300-2303113010231131-0200232233003001-2133332201101102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use management network.

Upstream description:

This can be used for messages where no values are needed.

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
use_management_network = {}
```

<a id="canonical-2222300311023313-1013131030210231-2111100120311200-0203331002230003-0011322122020022-2021223201332213-2322131123203003-0133131022112222"></a>

## Direct properties — use_management_network / 302312021211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101333020202312-1100122311322210-3030221331133122-0120212301313220-3201212030213131-2122121211110031-2200202302212100-1003033300022203"></a>

## Next pages — use_management_network / 302312021211 / 4

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0031021002303321-3112130012003112-1011131101211211-0033112221120300-3311120313031021-3210203333312121-1200221213301132-1033211302020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311122200321112-2101331231123320-1021212033200201-3120203223032121-3030323322021303-0213312222012022-0322123331132232-2230310333030222"></a>

## log_receiver_with_net.use_slo_sli — use_slo_sli / 200323220212 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030)
- log_receiver_with_net.use_slo_sli

<a id="canonical-1223031322103113-0031222033011113-3130120030202213-0321001022302333-0302331001120321-2323201302112202-3033011101232022-1300022001203033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use slo sli.

Upstream description:

This can be used for messages where no values are needed.

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
use_slo_sli = {}
```

<a id="canonical-2232200233232101-1133203212320033-0122230201331323-0012121031023213-2201012302302313-3321211023333211-3222230002100003-0102102202221233"></a>

## Direct properties — use_slo_sli / 200323220212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103322112103330-3011000112201302-3323031031232030-1013333332211000-3020203330213023-1122003320132310-3021123022320202-1120130330033022"></a>

## Next pages — use_slo_sli / 200323220212 / 4

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2210233313020231-0310131332331201-3031332230130222-0133101033233310-0223230331130213-2103112112102321-1230002312101232-1211013133212221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132113203021122-0301033012300320-3323323123130111-1020122230022101-0301021110212022-3030111213120212-1232332001301033-0131210323031120"></a>

## logs_streaming_disabled — logs_streaming_disabled / 011101210122 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- logs_streaming_disabled

<a id="canonical-0111001211331333-3333100033323212-2110200312112313-2301000232032002-0331312312200211-1233122002101233-2323212221312020-1333023010232331"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
logs_streaming_disabled = {}
```

<a id="canonical-2023222112130122-1301002113000220-1211131312123223-1020121020233322-3200012000212203-1033111313330310-1312233120322010-2321302003133030"></a>

## Direct properties — logs_streaming_disabled / 011101210122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131001010301221-3112023111111211-2013232123031233-2033333112213003-2310001212023032-1330322301001230-0012012002003100-2303302221230321"></a>

## Next pages — logs_streaming_disabled / 011101210122 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3231001003211033-1113201231132322-0013121220032323-2031203003230221-1133223030313302-3131212323231233-1301021322010201-1030310211331131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320321002320131-2210322010101110-1022320332311032-2112221330313110-3211233001301021-1102213113232332-1120113220111213-2210212032301132"></a>

## no_forward_proxy — no_forward_proxy / 020233210013 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_forward_proxy

<a id="canonical-1322030320001001-0330100333213223-0220321230132233-2313313023123323-2030302332103211-2033330212203300-3012021231012103-3230222113113031"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

Upstream description:

This can be used for messages where no values are needed.

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
no_forward_proxy = {}
```

<a id="canonical-0230223311333030-1332130230331000-0221111113020100-3212030002013030-2331011013203233-0220310301013220-0012221000321101-1221032212201122"></a>

## Direct properties — no_forward_proxy / 020233210013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002023030002332-1022302320123220-3113100012332132-2202001003102100-3132300003220121-2021201002303032-0001333030203011-0121102330312120"></a>

## Next pages — no_forward_proxy / 020233210013 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1111331220330301-2030301032232003-0321010121010110-1033012330011003-3321221320333313-2222101321233313-3201213303322110-3011030103300011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030120023103021-1330212200132210-0211303030332103-0322222100010100-3321103310223301-1103133112303033-3201103211022031-0120222201300233"></a>

## no_network_policy — no_network_policy / 103202023123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_network_policy

<a id="canonical-2213010322031120-2120333011333213-0322320013333003-3120023210232013-0222210010203123-1032320102313113-1233002003220122-1232110111321031"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

Upstream description:

This can be used for messages where no values are needed.

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
no_network_policy = {}
```

<a id="canonical-2113023321012123-1221333321221113-3212011030200023-2210120023313122-3210001211002121-1123313231212203-0331130111102232-1201200230300312"></a>

## Direct properties — no_network_policy / 103202023123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331000201223010-1231201012002232-2012211320320301-1032123323300310-2110231031003020-2101203020030300-3220032202333212-3100103231232133"></a>

## Next pages — no_network_policy / 103202023123 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3123331333122110-3122333032323313-1322031023033330-3022013201322102-1030322011123202-1032113110203010-1103213002302323-3011223220033322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220021333123111-0120203212013322-3010202103012111-1231011003111213-1311322230030001-1003001003323013-0322132121322300-1301031023201012"></a>

## no_proxy_bypass — no_proxy_bypass / 001011113000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_proxy_bypass

<a id="canonical-3022023002231032-3101230312322203-1210102303113020-3133121101220113-3332021101223101-0031033231310313-2221010131313222-1131033300021022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no proxy bypass.

Upstream description:

This can be used for messages where no values are needed.

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
no_proxy_bypass = {}
```

<a id="canonical-1100022233203300-1011102111322013-1003320121330221-3300033123130211-1113102203330233-3100101223221302-0133001033313223-2303210131233313"></a>

## Direct properties — no_proxy_bypass / 001011113000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211310032322333-3322011330232232-1033011111320203-3020013200100133-0321301202211020-0200320322131230-1002303203232330-3133133020013032"></a>

## Next pages — no_proxy_bypass / 001011113000 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3322012222010332-2230100331221211-0132203200323130-2003020232110130-0012210112022001-1220312313020321-1303222003211323-2020120133231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103331300000031-0110012321312211-1010000111120133-1031322030010132-2030103002002302-3330330330131000-2313002310202231-1212133122010101"></a>

## no_s2s_connectivity_sli — no_s2s_connectivity_sli / 200131132131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_s2s_connectivity_sli

<a id="canonical-0330303002222222-1213002233022103-0000311222000331-1230302330001022-0022102302021300-2031322212021101-0113011320113023-0212000011230312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no s2s connectivity sli.

Upstream description:

This can be used for messages where no values are needed.

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
no_s2s_connectivity_sli = {}
```

<a id="canonical-2023111303133102-2033313330031331-1002233103000111-0013220112123212-1113331210121330-2213213123022010-0201030111231300-0003300122301123"></a>

## Direct properties — no_s2s_connectivity_sli / 200131132131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232311001300113-1200133301232213-3302200130332202-2002230322321323-1123100130312102-0103031030201212-3300011021232202-3310231233111011"></a>

## Next pages — no_s2s_connectivity_sli / 200131132131 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1032112221231302-2123032103000023-1212321030120112-1301100133021333-3122002203103112-2022222012000331-0333232203322311-3201133120130202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320231321313100-1013310210131010-2203333323122301-2103000111330100-2113203000113101-2211302023321022-3212332101320202-1220033322120303"></a>

## no_s2s_connectivity_slo — no_s2s_connectivity_slo / 213310323130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- no_s2s_connectivity_slo

<a id="canonical-1030303003312203-3112321311031221-2310332332000120-1233313121311313-2111010213232133-2302111301212233-2211333200111312-1310302310103121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no s2s connectivity slo.

Upstream description:

This can be used for messages where no values are needed.

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
no_s2s_connectivity_slo = {}
```

<a id="canonical-0202213123232123-0203331021123230-2203323221221122-2000001303121300-2332201123333131-2323211201303012-1001130121222203-3033112312233103"></a>

## Direct properties — no_s2s_connectivity_slo / 213310323130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330000100123021-0000321021223111-3310103222112033-0000131023013323-1231002313022032-2100210120103112-1130112113030303-0111113131220100"></a>

## Next pages — no_s2s_connectivity_slo / 213310323130 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200312212032012-3130021302132010-0300101013310221-0221032101211321-0033313002021002-1001201330131221-1020011223233302-0310112103131312"></a>

## nutanix — nutanix / 003312323100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- nutanix

<a id="canonical-1210012123223033-3213322002010300-1021212010202320-1123300232210200-2332210120200203-3222221311223100-3020212001233310-2032320102120102"></a>

Type: `"object"`. single nested block, Optional.

Nutanix Provider Type. Nutanix Provider Type.

Upstream description:

Nutanix Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

Terraform syntax:

```terraform
nutanix {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002212213300310-0000131132213132-1202332133313100-2001011130332131-0231100121130120-0002330130000300-1212030030312021-2132002212023212"></a>

## Direct properties — nutanix / 003312323100 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000): complete subsection reference.

<a id="canonical-0033333022333033-0212201120112100-3310013220031033-0013203121131330-1121012101003103-0211022320222033-0122320303111011-3101121012231131"></a>

## Next pages — nutanix / 003312323100 / 4

- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233331103031300-2232110122122132-0001103231020001-3031223313300030-2020112103021133-3102311001332133-0222030203000323-2111120303230203"></a>

## nutanix.not_managed — not_managed / 201320220002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- nutanix.not_managed

<a id="canonical-3332103223333133-1331303223121321-0101200130020110-0123001130303122-1221300203223210-0121012031121122-2220132330303212-2000101111011231"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233111232200323-2231023110100210-1333103212131211-1223012223021001-3031200200311001-2120212231303022-0213001012103230-2212122310331023"></a>

## Direct properties — not_managed / 201320220002 / 3

- [node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231): complete subsection reference.

<a id="canonical-3010232023302223-0130021320031301-1202020332303113-2021112333303333-1202121311023332-0221303021111223-2201332132113201-2131013103022130"></a>

## Next pages — not_managed / 201320220002 / 4

- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220211213332313-0200322312313030-0131110201033211-3333023021112231-0223220121322232-1000110120312110-0002110320032333-3033211311101113"></a>

## nutanix.not_managed.node_list — node_list / 212110022330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- nutanix.not_managed.node_list

<a id="canonical-1210031322321012-3222110033003100-0212011023002023-1223010201003232-3023100121212113-0003211310331302-0320022220132331-1323310012111122"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020010121101231-0331300322213213-1300201231202302-1333022111200230-2210212323021310-1202011021101211-1303002113033332-3030000101200123"></a>

## Direct properties — node_list / 212110022330 / 3

<a id="canonical-3201222221111330-0132222103330002-0312230333203021-3101122003232112-3101020211112200-0121003003323111-1332001300303301-3233111121113132"></a>

<a id="canonical-3131123331222033-3121121222002200-2231003210232212-2323310001111331-2132131103221013-1321203203113011-0003030221123030-0203232230012232"></a>

## hostname property — node_list / 212110022330 / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202): complete subsection reference.

<a id="canonical-3121223103200210-0212302310101201-2232102210013313-1223031203010322-1100302110001311-2232332320312323-1031032233223000-0322223132302112"></a>

<a id="canonical-0233002333231310-3330021101000232-2222221212332212-2110201303000023-2322200233103112-3222023000132013-2230301332331203-0300133230011023"></a>

## public_ip property — node_list / 212110022330 / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0000300030011020-1301120123112320-0311033212112210-3322123122233112-2133113230113331-3330030211221203-2321333121221311-3331130021230033"></a>

<a id="canonical-0310231302021312-1030223112101200-2131100203223233-2232232030011101-2312102300222213-2101110120022012-3302321103330322-3001212023132233"></a>

## type property — node_list / 212110022330 / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-3311101003132012-3022122300332033-3020103301231101-3120323330302222-3212010011103211-1020101021213022-2333232330020133-0201111301313033"></a>

## Next pages — node_list / 212110022330 / 7

- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032113011120003-2120310100201323-3003101230020213-3230030220103220-3012223223022233-2121003023033302-3302303322333121-3322131201333011"></a>

## nutanix.not_managed.node_list.interface_list — interface_list / 331312311101 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- nutanix.not_managed.node_list.interface_list

<a id="canonical-2321302213113003-0000212013222013-1021312112322210-3000111201312312-0110101012210203-2121002310003332-2130312001200220-1121300212111112"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323303211311012-1021230221231113-1200122232303121-1003323232122232-0031223303302221-3301222130103003-3011023312230123-0331123200022022"></a>

## Direct properties — interface_list / 331312311101 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312): complete subsection reference.

<a id="canonical-0010312010130023-2323303130223323-2120301303230321-2323311000303202-2312000112002313-2103012212212300-2302131200203033-3210310300130102"></a>

<a id="canonical-2211301222010002-1123110301021203-1111333200003032-1302103320210213-2113001221112313-3021313023121110-2331131030113320-1122332032130131"></a>

## description_spec property — interface_list / 331312311101 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-012.md#canonical-0122301031222001-3012001030032031-1220230130013231-2113030122330033-3223230311003133-2330110130310303-1103123101201022-1002300311110013): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-1220303103223201-0002300332200330-0210220311211213-3220113033021232-1111312303002001-1030223331130300-3322301012231302-2331321131003323): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202): complete subsection reference.

<a id="canonical-3101313001320113-0130203231201300-2111300231102020-0232113013101313-2232013303121131-3022201100020333-1130112300212022-1020311332313332"></a>

<a id="canonical-1131033113303020-0130003222120200-1133311232233130-3113310333300033-3202300303220011-2332023202130132-1320323132213310-3213232102223023"></a>

## is_management property — interface_list / 331312311101 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-3130022230312230-1112220311321101-2031300230100311-1223031111000133-1021332001101030-2003133122202101-3221231230313211-0213012233331223"></a>

<a id="canonical-2321020030130312-3002113120232131-0121300030223111-1230311210121020-0312113032323002-3022322202211202-0222010010303001-0033320331212102"></a>

## is_primary property — interface_list / 331312311101 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1022332031010300-1100001212003330-1223231012002103-1220121111110222-2022301201200020-1203201023033020-1332131302223332-2221112132231222"></a>

<a id="canonical-0232210121011233-0120312010311012-2323103132100303-2322332000311100-0103133033002111-1030131022233101-2202032311003000-1310210333021230"></a>

## labels property — interface_list / 331312311101 / 7

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--reference--group-013.md#canonical-3123320032100001-1031133131223231-0011130301203201-1123122201301021-0120221213231131-1110333302111203-0032120301303121-3232123031302301): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-3123322001322102-3101003102311001-3021021212103320-1211232031000131-3021021232200112-1233221230310332-3303221022112021-3131022012122013): complete subsection reference.

<a id="canonical-0133112101030022-1231321122311313-3010232200323131-2211231323100212-2323111320330003-1321331012322130-3012010123022221-3023220320300220"></a>

<a id="canonical-2023111122131120-3233121233330310-1021333231311020-3013321333201332-3122033130320103-2113032113131022-0120231013301310-3313310112331320"></a>

## mtu property — interface_list / 331312311101 / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-3303313332011120-0302110011333230-2300232012011022-0311202333023312-3101323312221331-0020113303011231-2313011333332231-0300002010120013"></a>

<a id="canonical-1130011301222220-2230021332002001-2301322000122010-3321321101311200-2330030222130203-2131100002320133-0113000333311311-2323032132133312"></a>

## name property — interface_list / 331312311101 / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--reference--group-013.md#canonical-2033233301013022-2201212323332122-2311010013222211-0233122301232213-2022203031021300-2002200331012323-3201033130200012-2133230201113131): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-013.md#canonical-1300113120101223-1220103101320303-2033212212031203-0302232302113211-3320033201230113-3331203111100003-1222310233300013-2003010201023320): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-1003012130113301-0102301003301331-3223303121002030-1103313212120000-3101301301110332-1210200200131110-2221021033010310-1313302111203121): complete subsection reference.

<a id="canonical-0022222311123200-3301203111032300-1332003100313322-3002011112111303-1021110313320033-3133122020222223-0232330131031112-1201011230213231"></a>

<a id="canonical-1033003121333223-1112300010212032-0210102001211230-3212122301031012-3011303201013131-3011010221220333-1133313213111210-2132030123302022"></a>

## priority property — interface_list / 331312311101 / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-1231301312010313-2213232000112222-3013230012232221-0133121113311310-1103002210331213-3120322321210302-3322223321233130-2113230112301313): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-013.md#canonical-3213332311302012-0322110013322032-2023302321211111-0301121010312031-3201102121211013-0110331013133311-3000110302320313-1333330332032231): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-013.md#canonical-3000311211321033-1020312111303303-0333321101203202-0120301212302320-2133131020320003-0302021131332021-1303313201120311-0232213200222202): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-2132202020202332-1110010011212323-2210113032102032-1032011021313132-2320203033333212-1130002223312202-1300201210000223-0101311100130333): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-0001022032320011-2010231101113231-3102012231032321-2010221100201020-2131313312332202-3123322133001323-3312023211310120-1100230201223210): complete subsection reference.

<a id="canonical-2201222333103212-2232211302103121-1102221000333012-3231110010120002-3231200103112130-1311033022232130-1100123131120220-0323221320001321"></a>

## Next pages — interface_list / 331312311101 / 11

- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312)
- [nutanix.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-012.md#canonical-0122301031222001-3012001030032031-1220230130013231-2113030122330033-3223230311003133-2330110130310303-1103123101201022-1002300311110013)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220)
- [nutanix.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-1220303103223201-0002300332200330-0210220311211213-3220113033021232-1111312303002001-1030223331130300-3322301012231302-2331321131003323)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-013.md#canonical-3123320032100001-1031133131223231-0011130301203201-1123122201301021-0120221213231131-1110333302111203-0032120301303121-3232123031302301)
- [nutanix.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-3123322001322102-3101003102311001-3021021212103320-1211232031000131-3021021232200112-1233221230310332-3303221022112021-3131022012122013)
- [nutanix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-013.md#canonical-2033233301013022-2201212323332122-2311010013222211-0233122301232213-2022203031021300-2002200331012323-3201033130200012-2133230201113131)
- [nutanix.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-013.md#canonical-1300113120101223-1220103101320303-2033212212031203-0302232302113211-3320033201230113-3331203111100003-1222310233300013-2003010201023320)
- [nutanix.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-1003012130113301-0102301003301331-3223303121002030-1103313212120000-3101301301110332-1210200200131110-2221021033010310-1313302111203121)
- [nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-1231301312010313-2213232000112222-3013230012232221-0133121113311310-1103002210331213-3120322321210302-3322223321233130-2113230112301313)
- [nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-013.md#canonical-3213332311302012-0322110013322032-2023302321211111-0301121010312031-3201102121211013-0110331013133311-3000110302320313-1333330332032231)
- [nutanix.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-013.md#canonical-3000311211321033-1020312111303303-0333321101203202-0120301212302320-2133131020320003-0302021131332021-1303313201120311-0232213200222202)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-2132202020202332-1110010011212323-2210113032102032-1032011021313132-2320203033333212-1130002223312202-1300201210000223-0101311100130333)
- [nutanix.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-0001022032320011-2010231101113231-3102012231032321-2010221100201020-2131313312332202-3123322133001323-3312023211310120-1100230201223210)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301223032201121-0002310221202000-0221223113132231-2010132031333230-0300331221332121-1330013313333202-1020111111300232-2103011303111011"></a>

## nutanix.not_managed.node_list.interface_list.bond_interface — bond_interface / 322303320111 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.bond_interface

<a id="canonical-3120123122112122-0313131101112011-2101302303032130-3032233013101222-2121231200313323-1313100220000213-2102100300023032-2333303122233321"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212200312302111-3110000001302021-2232303331320213-3022320103221030-0303011021330300-2112130011021031-3012021003333110-2311212213113212"></a>

## Direct properties — bond_interface / 322303320111 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-012.md#canonical-3111202113131003-0213120221131113-1123323300322122-3121230102333222-0110133110311013-0230210303310310-1220030112120023-3012013333112330): complete subsection reference.

<a id="canonical-3102033031033303-0210300123301023-2200023023221303-0201013003013021-0303223130133013-3002011130101023-2200021220231101-3223300320311300"></a>

<a id="canonical-2330211303310001-0230133020211100-3232320031023320-0213200302221312-2110301000201010-1100330123212320-3320013131131001-0100112321033320"></a>

## devices property — bond_interface / 322303320111 / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--securemesh_site_v2--reference--group-012.md#canonical-2202133101000121-1033000230102033-2013100322323133-2013101102103103-0212203013023001-1300133001031230-1030131212302111-2130322331331120): complete subsection reference.

<a id="canonical-2310103201222112-2001212110232333-1021302003002121-2030021233210021-0121220030201123-2222130212223023-3032331001310131-2211131013113023"></a>

<a id="canonical-3312130112112120-0030010010321323-2033312323222031-2221122321332203-3002221012230122-3003331011330030-1111033102020330-0121333112123111"></a>

## link_polling_interval property — bond_interface / 322303320111 / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-2232332202030010-3200133112332303-3101330132122002-2301123330033331-2031032222121131-3113132212332301-2321013300130102-1310011332002120"></a>

<a id="canonical-3122032110321323-2013123312331330-1002213212322233-2021122021021000-3202132313000332-2310202131302011-0010322311103323-0122221022131230"></a>

## link_up_delay property — bond_interface / 322303320111 / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-0210133202120311-0322202310012331-2333100201220230-0331201100333033-2223320123130322-1123323310200230-2002210000033012-2311001030030122"></a>

<a id="canonical-2033010310123323-1112123212023133-3232210203333211-2112320311002123-2311031333210200-0033020030113112-2102021312020210-3131132300301022"></a>

## name property — bond_interface / 322303320111 / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1322222101222213-1032021320121233-3200202302202120-3101321323101231-0302132131131330-0211030010320213-3222223112120133-1322310322230223"></a>

## Next pages — bond_interface / 322303320111 / 8

- [nutanix.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-012.md#canonical-3111202113131003-0213120221131113-1123323300322122-3121230102333222-0110133110311013-0230210303310310-1220030112120023-3012013333112330)
- [nutanix.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-012.md#canonical-2202133101000121-1033000230102033-2013100322323133-2013101102103103-0212203013023001-1300133001031230-1030131212302111-2130322331331120)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3111202113131003-0213120221131113-1123323300322122-3121230102333222-0110133110311013-0230210303310310-1220030112120023-3012013333112330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032210213130202-0322223321201031-2200012312102210-1123012132223220-2221320310210103-2012223220121120-2212112002323323-2130000013323213"></a>

## nutanix.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 310321122002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312)
- nutanix.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-2123310022220210-2311213321320030-2002031012322313-3103223231320013-2030303032130211-0023110331022333-0032210232230201-0001202332012131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

Upstream description:

This can be used for messages where no values are needed.

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
active_backup = {}
```

<a id="canonical-0122120030121221-1013123312213312-2332310202233333-0222032032121222-3322312101002230-0113222110332120-3322302333122330-2301001003023200"></a>

## Direct properties — active_backup / 310321122002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203223302330300-0212013231223232-2003331112311020-1330231310132100-0111323331032002-0202233211100333-0131213331133203-2120111113123010"></a>

## Next pages — active_backup / 310321122002 / 4

- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2202133101000121-1033000230102033-2013100322323133-2013101102103103-0212203013023001-1300133001031230-1030131212302111-2130322331331120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332120212010133-3302311220012003-3133123231022022-3103001031303333-1232130312311031-0222233033030132-3311032312131021-2200022002300023"></a>

## nutanix.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 122113130000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312)
- nutanix.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-0030210333102012-1311331330203212-0121330121302300-3211133122303003-2200331302223133-1030320001322310-3122113002020220-1301112332321131"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212230220010100-0130321303121220-3311301031131312-3103131320011021-0013112010121301-2010220020211300-0300102133311112-2020231020222321"></a>

## Direct properties — lacp / 122113130000 / 3

<a id="canonical-2013021322033013-0330223203333131-1331223120311122-1003323100102200-2321000110100322-0301311323223103-3213103021202231-3112013102212211"></a>

<a id="canonical-1322101310001303-2030131223212230-3113110313133301-1203001331203212-3302123222001032-1332032200210221-3303011332201313-1011300332011013"></a>

## rate property — lacp / 122113130000 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-2210321020031010-3132202022301111-1021222200223330-3103132000331210-2202200110031133-2032101002120310-2303213313022310-3333021022032210"></a>

## Next pages — lacp / 122113130000 / 5

- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0122301031222001-3012001030032031-1220230130013231-2113030122330033-3223230311003133-2330110130310303-1103123101201022-1002300311110013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002223300311223-2103321320221011-1223022110110232-1232012012131021-3100112301023301-0021032203221211-2312120210200311-0211011220231332"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 033103202023 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-0211122002310120-0102301312113310-2200322113210212-1211313030030231-2212203231120102-1021033202323131-0303231200123123-3121102112032130"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
dhcp_client = {}
```

<a id="canonical-3012322310332202-1232033303300001-3331322301123113-1002103103200333-0233032000323322-1103132202101120-1312100000011123-1223131313313020"></a>

## Direct properties — dhcp_client / 033103202023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001230033131312-0033220210023310-3210222333031021-1123102001100320-3020230331301300-2201302320312011-2233020121113102-0300023021031310"></a>

## Next pages — dhcp_client / 033103202023 / 4

- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102121221013113-2323300111122012-3110022202102222-2101020233103320-3123311313122113-1103232232233123-3312301021222112-2132311120120202"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 013021111023 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-0130003112301020-0220201333202102-0103312132301330-3001103332321202-3103021231203002-3131121111110232-3031131211113201-3110322312333113"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003012001313300-3302233120132111-2302023212012322-3033000101203101-1202121332003203-1012332003231022-1210101020203320-2021101101023220"></a>

## Direct properties — dhcp_server / 013021111023 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-013.md#canonical-0121011231312233-2213310221211130-1211000322220132-3333011011201332-2332321003102103-1223130233032031-0330222100332000-0020323222230230): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-013.md#canonical-0102023312223300-0033330111301323-0113231113102132-3232212102023131-0033000231001121-0312332033113330-1300201012313011-3232020302002331): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-0310311333332031-1101221222033203-1111222202122001-3321011303221201-2111130230122020-0230100321321010-3221130311130223-1101310210023101): complete subsection reference.

<a id="canonical-0321330132101310-1002321323222223-0020203200133211-1023113030303220-2012221022322103-0000210000321330-0310230133313021-2211202032220231"></a>

<a id="canonical-1011023030221301-3310300211310310-1222033303203023-3013320111332003-1111100213022332-1010201130111023-2022200122112333-3001010133122122"></a>

## dhcp_option82_tag property — dhcp_server / 013021111023 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-2220033023202202-2122330031011220-0332131033221311-0021103021332013-0233200310033321-1111100330030330-2323100021212331-2221210010201333"></a>
