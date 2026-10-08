---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3131030030231103-3100013311313213-2312122111023223-1332210200121031-0330203011101002-2022022223113302-3031131031332103-1112001030211132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.blocked_service.dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [blocked_services](resources--securemesh_site_v2--reference--group-005.md#canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122)
- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-005.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- blocked_services.blocked_service.DNS

<a id="canonical-1313133131030320-1310101202031021-2323201013102223-1101220122201111-0031010332323001-2311001211030200-3232332222111122-0213033220132212"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
dns = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013313321133020-3000322300332212-2023330213032300-1101011303221033-0003302332133321-2133300013200133-1031111333331013-0230020002023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.blocked_service.ssh` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [blocked_services](resources--securemesh_site_v2--reference--group-005.md#canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122)
- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-005.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- blocked_services.blocked_service.SSH

<a id="canonical-0333302002332220-2132030323021223-1301110310321313-2110121232033303-2002113030200230-2113313030002011-2313231000122210-2121111030201112"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
ssh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311113333110132-1332331102333301-1033301322332030-2201030230223030-1032031310230003-2320012320231200-1301031130212110-1031030131310223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_services.blocked_service.web_user_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [blocked_services](resources--securemesh_site_v2--reference--group-005.md#canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122)
- [blocked_services.blocked_service](resources--securemesh_site_v2--reference--group-005.md#canonical-3311110301202031-3012023312233113-3323122301011333-1110311100211111-0012323011332303-0033312120030231-0310110101333113-1223301211202032)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-1331200200130100-0002102132110323-2011012011321222-2333332031310103-1000022000122032-2033303011202321-3231220110301310-2202321210013111"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
web_user_interface = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_proxy` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- custom_proxy

<a id="canonical-0133132100102112-1321232211122100-3231121222211020-1323020332320220-3100221323323313-2002012322321002-1030330013212022-1310232101232330"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_proxy, f5\_proxy, private\_adn\] Configuration parameter for custom proxy.

Additional upstream details:

Custom Enterprise Proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("proxy_ip_address",
    "proxy_port"),
  validators.ConflictingObjectAttributes("disable_re_tunnel",
    "enable_re_tunnel")}
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
  "x-ves-oneof-field-use_for_re_tunnel_choice": "[\"disable_re_tunnel\",\"enable_re_tunnel\"]"
}
```

OneOf alternatives in this subsection:

- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-0133132100102112-1321232211122100-3231121222211020-1323020332320220-3100221323323313-2002012322321002-1030330013212022-1310232101232330)
- [f5_proxy](resources--securemesh_site_v2--reference--group-009.md#canonical-0323202033202102-2203130323103033-2120333100022103-1031333331023101-0001233230123023-0302313331331233-0112221333211032-1203223313013302)
- [private_adn](resources--securemesh_site_v2--reference--group-016.md#canonical-3001100011332303-2231203223000211-1230200300233313-2012022223033331-1002001201232231-1312011231123311-2020311101200103-0010032010320220)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212133223201312-1032302310112133-2111023110131021-2031310312232213-1203212010032103-3333032001233011-3333220231000032-3022312111303100"></a>

### Direct properties for `custom_proxy`

- [disable_re_tunnel](resources--securemesh_site_v2--reference--group-006.md#canonical-2211130001212222-0222130033220113-2331102010101102-1233031020203303-0233103111010211-2321212021332000-1232110123223311-2232011002301213): complete subsection reference.

- [enable_re_tunnel](resources--securemesh_site_v2--reference--group-006.md#canonical-3111102123022002-1031033320333131-2212221101213001-2311200122113001-1132310110311032-0212310303232030-2022223013213220-2022110121020030): complete subsection reference.

- [password](resources--securemesh_site_v2--reference--group-006.md#canonical-3001030210312201-2020010031322113-3001013121310303-2030021121311010-1300301111130220-0312313102131210-3231122222311300-2113320131233010): complete subsection reference.

<a id="canonical-0020010033100320-0210101320112332-1100201312321323-2212132123303213-2101112333232013-1232201123033220-3021211221013003-2033110213333020"></a>

<a id="canonical-3203231201022130-1300022033213003-3331012033330023-2233100203120300-3223011130221201-0003121220013233-1232112230313101-2133322330101320"></a>

#### `custom_proxy.proxy_ip_address` property

Type: `"string"`. Optional.

Specify the IPv4 Address of the internal Enterprise Proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1003102021201221-2122201000221001-1230301333000122-2100000003000032-1130233323330221-0330202123023113-2230032100020132-0021202303321113"></a>

<a id="canonical-2123011022030230-1002032333122321-2332323132233222-2102123310012123-0300132101023111-1020212000310122-3102130303311212-2202013203112131"></a>

#### `custom_proxy.proxy_port` property

Type: `"number"`. Optional.

Specify the Port of the internal Enterprise Proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0130213202312113-2000012122332021-1303013322111301-2001202112302232-3012213023300212-3123230323231110-2033012300102123-3003101213332232"></a>

<a id="canonical-0131103121032220-2032102302211022-0033311030133203-0010311010010023-1333210030111120-1231032112001222-1121120032210011-1033003200133333"></a>

#### `custom_proxy.username` property

Type: `"string"`. Optional.

If the internal Enterprise Proxy is using basic authentication, specify the username. This is an
optional field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "identity",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2211130001212222-0222130033220113-2331102010101102-1233031020203303-0233103111010211-2321212021332000-1232110123223311-2232011002301213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_proxy.disable_re_tunnel` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131)
- custom_proxy.disable_re_tunnel

<a id="canonical-0222132233202311-1201122313011312-3033332201320120-2321223013222012-3331221311313002-2101233132312200-1221031102232311-3322032120321012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable re tunnel.

Additional upstream details:

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
disable_re_tunnel = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111102123022002-1031033320333131-2212221101213001-2311200122113001-1132310110311032-0212310303232030-2022223013213220-2022110121020030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_proxy.enable_re_tunnel` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131)
- custom_proxy.enable_re_tunnel

<a id="canonical-0323330233113230-2322321011003320-1002323332210003-3322233232000123-2120102000232021-1201100020311221-0000020133030031-3001222002031021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable re tunnel.

Additional upstream details:

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
enable_re_tunnel = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001030210312201-2020010031322113-3001013121310303-2030021121311010-1300301111130220-0312313102131210-3231122222311300-2113320131233010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_proxy.password` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131)
- custom_proxy.password

<a id="canonical-3033312022101300-0232210302300020-2013310103202310-1212300333131323-2323101302312300-1312130002102011-1021020222120212-0302301313000332"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131112203030133-0303020211023313-3220033200310020-0022010111323113-2313321133120201-1003032002203312-1200030100321333-3100321221313330"></a>

### Direct properties for `custom_proxy.password`

- [blindfold_secret_info](resources--securemesh_site_v2--reference--group-006.md#canonical-3210331201211223-3002113002300322-0221113212310013-0133013303230121-1013301122021020-0223003300003232-3023321202300220-1011223111211320): complete subsection reference.

- [clear_secret_info](resources--securemesh_site_v2--reference--group-006.md#canonical-3223030303103201-2313210133333313-1120221112033100-1332302313313131-2023322120302103-1310021221110300-2130010231132310-0131332332031013): complete subsection reference.

<a id="canonical-3210331201211223-3002113002300322-0221113212310013-0133013303230121-1013301122021020-0223003300003232-3023321202300220-1011223111211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_proxy.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131)
- [custom_proxy.password](resources--securemesh_site_v2--reference--group-006.md#canonical-3001030210312201-2020010031322113-3001013121310303-2030021121311010-1300301111130220-0312313102131210-3231122222311300-2113320131233010)
- custom_proxy.password.blindfold_secret_info

<a id="canonical-3213132010110031-3033323101123123-0203300322101322-3200012000223111-3030033331222011-2311112033123332-1322231231121032-3313322323033221"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121121110102310-2222113010011120-2312032302103031-2203320110212223-0032322010133001-0313121111213220-1002313311021222-1133302133332312"></a>

### Direct properties for `custom_proxy.password.blindfold_secret_info`

<a id="canonical-2113201021003313-2023323221131131-1020131300122022-2111231030122210-2321022033030113-1120030020303302-2121012223203221-1011223032011002"></a>

#### `custom_proxy.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0000012230300123-0101333033030111-2232203222223233-1032202120200330-3220120112101131-0130311131212120-3332113011201301-1230012001103001"></a>

<a id="canonical-1130113320022310-1320001022322312-2022212112011132-0032111302313110-3010230102213321-1221201101033100-0103211210033211-1132321020120302"></a>

#### `custom_proxy.password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3011321001133133-0033201011210002-2332303212111201-1322121113031212-2131100112003313-3320112333203001-3121102200032021-3312203323120301"></a>

<a id="canonical-3122312311203312-0030333013231221-3300311220222301-2220132311021213-0203231310030120-2232200012001231-1011311220031112-1121111200202310"></a>

#### `custom_proxy.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3223030303103201-2313210133333313-1120221112033100-1332302313313131-2023322120302103-1310021221110300-2130010231132310-0131332332031013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_proxy.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131)
- [custom_proxy.password](resources--securemesh_site_v2--reference--group-006.md#canonical-3001030210312201-2020010031322113-3001013121310303-2030021121311010-1300301111130220-0312313102131210-3231122222311300-2113320131233010)
- custom_proxy.password.clear_secret_info

<a id="canonical-1232210121101021-0213110130101122-2232202030310230-3330330020113033-3333312333033232-0213312133100203-1131003031221312-0322022320132023"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021012112202220-3311323222220021-3303301311002231-0212031330331030-1012230321121210-2230021001322213-2131323202312123-2321220223020233"></a>

### Direct properties for `custom_proxy.password.clear_secret_info`

<a id="canonical-3102321103020112-1011031211001022-3201332310323320-2100203300101023-0301330121020320-1230002003003313-1202333233200133-3131332222333111"></a>

#### `custom_proxy.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0223032110210203-1301020112322212-0332131212310112-1313313003301311-3100132332100303-1210001023111112-2131012202122012-2013203221020103"></a>

<a id="canonical-3020210031210120-3211013331111123-1012220102202031-0101010022123203-3020220122212133-1321022221100220-1311221022331220-2101103203033002"></a>

#### `custom_proxy.password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1201120122021230-3013113003303010-0312222210310120-1032233221132010-0003011010201111-3212223322123210-1233330103221220-3111220102212230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_proxy_bypass` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- custom_proxy_bypass

<a id="canonical-0321121310031302-3001323333210001-0332113021331022-3202330233302121-1322302022300111-3020133013030000-1333032113321213-2300231103033123"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_proxy\_bypass, no\_proxy\_bypass; Default: no\_proxy\_bypass\] Configuration
parameter for custom proxy bypass.

Additional upstream details:

List of domains to bypass the proxy.

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

OneOf alternatives in this subsection:

- [custom_proxy_bypass](resources--securemesh_site_v2--reference--group-006.md#canonical-0321121310031302-3001323333210001-0332113021331022-3202330233302121-1322302022300111-3020133013030000-1333032113321213-2300231103033123)
- [no_proxy_bypass](resources--securemesh_site_v2--reference--group-011.md#canonical-3022023002231032-3101230312322203-1210102303113020-3133121101220113-3332021101223101-0031033231310313-2221010131313222-1131033300021022)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_proxy_bypass {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112131010001000-3033221133301010-3322111110120103-3320120122100201-0231100003201213-3200332233122021-1301230313123130-0200131112320201"></a>

### Direct properties for `custom_proxy_bypass`

<a id="canonical-1133113202202210-3220332300222203-0103312001213200-2310121320113322-3022312001302021-3120002212002331-0231013010032012-2011102213323313"></a>

#### `custom_proxy_bypass.proxy_bypass` property

Type: `["list", "string"]`. Optional.

Proxy Bypass. List of domains to bypass the proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.hostname_or_ip": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname_or_ip": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1020321133311233-2110133311213121-2123333100012012-3200301200123213-1020103130321222-1021002013333211-1002020122222102-3103333110002323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dc_cluster_group_sli` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- dc_cluster_group_sli

<a id="canonical-1020300212301120-1320101102310232-1010130313301102-3202331322202133-1101131123210202-1101022221021110-0300103310332221-1303203313130333"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dc\_cluster\_group\_sli, no\_s2s\_connectivity\_sli; Default: no\_s2s\_connectivity\_sli\]
Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

OneOf alternatives in this subsection:

- [dc_cluster_group_sli](resources--securemesh_site_v2--reference--group-006.md#canonical-1020300212301120-1320101102310232-1010130313301102-3202331322202133-1101131123210202-1101022221021110-0300103310332221-1303203313130333)
- [no_s2s_connectivity_sli](resources--securemesh_site_v2--reference--group-011.md#canonical-0330303002222222-1213002233022103-0000311222000331-1230302330001022-0022102302021300-2031322212021101-0113011320113023-0212000011230312)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dc_cluster_group_sli {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133332103021202-3113322212201121-3120122323302031-3110011130220331-3200210330312233-3322320012112320-0231213001012300-1332313113123133"></a>

### Direct properties for `dc_cluster_group_sli`

<a id="canonical-1211122111111332-3221001220320302-3213333311000133-2310131333003312-3303000222022000-0000211312311312-2303320022012233-2313111312232323"></a>

#### `dc_cluster_group_sli.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1031313320003220-1201021202120131-0033123132102103-3123003122012033-1312103330221301-2332121302323213-1002003220101022-0132211312221030"></a>

<a id="canonical-3312310313321031-0202000302233013-0312221301202003-2201100310310321-2322300020233211-2022103020203123-0301011333210232-1020012220333232"></a>

#### `dc_cluster_group_sli.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3133103022012023-0031122133021303-1033322023302320-0312323030100222-3322010301333023-2112320210300031-0202203213310213-3221030202001211"></a>

<a id="canonical-3012022013303003-0011021023201302-3031213321030313-2133123003312032-3313331023312023-3021132300021102-0333203102121011-0210002300331123"></a>

#### `dc_cluster_group_sli.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3100323033003130-1302331203022103-1233101210123010-3220222322211011-0332110232310011-1003002331302211-2010313222320033-2132103002200131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dc_cluster_group_slo` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- dc_cluster_group_slo

<a id="canonical-3120230010200133-0213301102030000-1103030023223132-1122222023103023-1312200222303103-3312320310320112-2131211200103212-1021200322100033"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dc\_cluster\_group\_slo, no\_s2s\_connectivity\_slo, site\_mesh\_group\_on\_slo; Default:
no\_s2s\_connectivity\_slo\] Type establishes a direct reference from one object(the referrer) to
another(the referred). Such a reference is in form of tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

OneOf alternatives in this subsection:

- [dc_cluster_group_slo](resources--securemesh_site_v2--reference--group-006.md#canonical-3120230010200133-0213301102030000-1103030023223132-1122222023103023-1312200222303103-3312320310320112-2131211200103212-1021200322100033)
- [no_s2s_connectivity_slo](resources--securemesh_site_v2--reference--group-011.md#canonical-1030303003312203-3112321311031221-2310332332000120-1233313121311313-2111010213232133-2302111301212233-2211333200111312-1310302310103121)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-3212222333120111-2230203121211023-0313133130013232-2310301323102000-1222112001023100-3302312232123100-3211022311121112-3323310000112231)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dc_cluster_group_slo {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232033113111320-1011133201032110-2012312231223321-2200333230233332-1230202100233333-0121323211120210-3223333321002302-0213221002333231"></a>

### Direct properties for `dc_cluster_group_slo`

<a id="canonical-1302220003223223-2222313131020001-0103202120003031-2321110213333020-2321133011101020-2313101210022103-1012201123030121-0022132312101322"></a>

#### `dc_cluster_group_slo.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1310300103011313-2321012221001100-1113330113131202-0320101301002321-3332210102200110-3030231313320121-2033301122200100-2022231310023322"></a>

<a id="canonical-0210021333201121-2302003101010111-3320230313233023-2210030120102212-0103001211032011-2131310330100003-2010311112330301-2211011021120332"></a>

#### `dc_cluster_group_slo.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0231101301113003-3101022323111210-0031132310312132-1321013231010303-3100313210033311-3213000212210230-3201333231321010-1003322112122232"></a>

<a id="canonical-0230332013233331-0022120010222011-1123032322312113-3010200331102022-3223322303013110-1212222232101013-0330122220223332-1221210302310200"></a>

#### `dc_cluster_group_slo.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1112203010312210-0132130311102202-0022222202032000-2131322332010323-3210102110223000-3200001233331101-1223311301211212-3200100111113210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_advanced_delivery` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_advanced_delivery

<a id="canonical-2122201323223101-2113232031003001-2120123002330002-2031202123102313-0110121120013130-3201031113231212-0023001220102221-0001120132122013"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_advanced\_delivery, enable\_advanced\_delivery; Default:
disable\_advanced\_delivery\] Configuration parameter for disable advanced delivery.

Additional upstream details:

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

OneOf alternatives in this subsection:

- [disable_advanced_delivery](resources--securemesh_site_v2--reference--group-006.md#canonical-2122201323223101-2113232031003001-2120123002330002-2031202123102313-0110121120013130-3201031113231212-0023001220102221-0001120132122013)
- [enable_advanced_delivery](resources--securemesh_site_v2--reference--group-008.md#canonical-1220131302000301-1122033110230313-0322113111321333-3220130110202322-2023020022301130-2201202233132330-3211202131020111-2200131101320320)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_advanced_delivery = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231002302323332-2013010223311332-0021120301210112-2232331120122110-0301210303200310-2002220333321202-2022103131122023-3101003100301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ha` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_ha

<a id="canonical-2320203132102000-0303312322222212-1301210010311110-2023032100003222-1021311001131211-3210210003302032-1322200012330113-2132221002021220"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ha, enable\_ha; Default: disable\_ha\] Enable this option

Additional upstream details:

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

OneOf alternatives in this subsection:

- [disable_ha](resources--securemesh_site_v2--reference--group-006.md#canonical-2320203132102000-0303312322222212-1301210010311110-2023032100003222-1021311001131211-3210210003302032-1322200012330113-2132221002021220)
- [enable_ha](resources--securemesh_site_v2--reference--group-008.md#canonical-2330001010130310-3032020220122111-2022223300332103-0310000223001331-1202013231320033-2100001200103323-1020221333133201-2231302313002023)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ha = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203112302132303-1313202311322332-0222023201233202-3221021303032331-1111112120320030-1100333003113200-0133320331321111-1203121122310001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_log_anonymization` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_log_anonymization

<a id="canonical-0103233332011021-3130012122032333-3333010121032130-3010112020123223-2103103122113023-1013312110211302-2223333230330322-1031302213200033"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_log\_anonymization, enable\_log\_anonymization; Default:
disable\_log\_anonymization\] Configuration parameter for disable log anonymization.

Additional upstream details:

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

OneOf alternatives in this subsection:

- [disable_log_anonymization](resources--securemesh_site_v2--reference--group-006.md#canonical-0103233332011021-3130012122032333-3333010121032130-3010112020123223-2103103122113023-1013312110211302-2223333230330322-1031302213200033)
- [enable_log_anonymization](resources--securemesh_site_v2--reference--group-008.md#canonical-1132023020201001-3210133312102302-2003130003203131-1011133331132330-0002300031323000-3311011021302323-3122332102101131-1011013211131323)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_log_anonymization = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131220221232201-2023022102131221-2033023020002031-0101312232133333-3212110223303200-0212322023231221-2122222012123200-2002031123021212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_management_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_management_network

<a id="canonical-0123200302122213-3223003012220322-1003113202213003-1331312133130231-1100212321202101-0000102330231113-2230230101320011-3133303132010323"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_management\_network, enable\_management\_network; Default:
disable\_management\_network\] Configuration parameter for disable management network.

Additional upstream details:

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

OneOf alternatives in this subsection:

- [disable_management_network](resources--securemesh_site_v2--reference--group-006.md#canonical-0123200302122213-3223003012220322-1003113202213003-1331312133130231-1100212321202101-0000102330231113-2230230101320011-3133303132010323)
- [enable_management_network](resources--securemesh_site_v2--reference--group-008.md#canonical-0312212202231003-0331322130101202-3220012301310333-1330003212202002-3111120113022031-3333231233121120-0000311222301201-3130230321003022)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_management_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302003210031020-0002223310123233-0230133013212201-1212020012311333-0112301130131323-0032203021212031-1033011301012302-3123112120202231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_url_categorization` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_url_categorization

<a id="canonical-3211111131131233-0121311023002102-0103031231100113-0101230013202231-2220222302212232-3111201103133120-2012213203113232-3210330222130300"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_url\_categorization, enable\_url\_categorization; Default:
disable\_url\_categorization\] Enable this option

Additional upstream details:

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

OneOf alternatives in this subsection:

- [disable_url_categorization](resources--securemesh_site_v2--reference--group-006.md#canonical-3211111131131233-0121311023002102-0103031231100113-0101230013202231-2220222302212232-3111201103133120-2012213203113232-3210330222130300)
- [enable_url_categorization](resources--securemesh_site_v2--reference--group-008.md#canonical-0322300022333021-3031321220220003-3121001123003331-3210300311220231-2113221133000010-1111302322003320-3311113331332300-0113223000130223)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_url_categorization = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dns_ntp_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- dns_ntp_config

<a id="canonical-0223123033322020-3331101211232032-2003212113133320-2330131311332313-3002122310030211-3302232302212222-2020003121330012-1010330002011332"></a>

Type: `"object"`. single nested block, Optional.

Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_dns",
    "f5_dns_default"),
  validators.ConflictingObjectAttributes("custom_ntp",
    "f5_ntp_default")}
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
  "x-ves-oneof-field-dns_server_choice": "[\"custom_dns\",\"f5_dns_default\"]",
  "x-ves-oneof-field-ntp_server_choice": "[\"custom_ntp\",\"f5_ntp_default\"]"
}
```

Terraform syntax:

```terraform
dns_ntp_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023213300122130-1133322203203031-2031222030032120-1122332122310111-1212211122302121-2321230221200011-0312221022333302-0232122011231313"></a>

### Direct properties for `dns_ntp_config`

- [custom_dns](resources--securemesh_site_v2--reference--group-006.md#canonical-3012001213011312-2021230120013323-2020301303130133-2103101312332311-3022013310032001-1213000122300303-1311133100123222-3132220000230031): complete subsection reference.

- [custom_ntp](resources--securemesh_site_v2--reference--group-006.md#canonical-0223311330322023-1031030023000003-2101020110023233-2222130102033223-2332103012302011-0020210112211022-1220232303021232-3032220330030110): complete subsection reference.

- [f5_dns_default](resources--securemesh_site_v2--reference--group-006.md#canonical-3013002232311000-0033131110202020-1213021032031033-1303213220200223-2332212011310211-1133223021130120-3333210012011000-2002210132030321): complete subsection reference.

- [f5_ntp_default](resources--securemesh_site_v2--reference--group-006.md#canonical-1013321102112123-3101002330233212-0213101211010332-1323311033231203-1022110213210023-1033323013110330-0000013023222121-0330331202122223): complete subsection reference.

<a id="canonical-3012001213011312-2021230120013323-2020301303130133-2103101312332311-3022013310032001-1213000122300303-1311133100123222-3132220000230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dns_ntp_config.custom_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [dns_ntp_config](resources--securemesh_site_v2--reference--group-006.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- dns_ntp_config.custom_dns

<a id="canonical-2112010003113112-1023301323001130-3121213111201030-2013130203000301-3002030020010112-0022102210132012-3222102032112100-3121210102210231"></a>

Type: `"object"`. single nested block, Optional.

DNS Servers. DNS Servers.

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
custom_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102302121101023-3132303313323112-3303331310102102-0002120303202222-1122031232223303-1013213303121202-2112131331131101-1121230003213330"></a>

### Direct properties for `dns_ntp_config.custom_dns`

<a id="canonical-0221120211322322-2123130332010213-2323313322210120-2130022031301112-0123213332000113-2200333200322030-3001310103011221-0010002132012130"></a>

#### `dns_ntp_config.custom_dns.dns_servers` property

Type: `["list", "string"]`. Optional.

DNS Servers. DNS Servers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0223311330322023-1031030023000003-2101020110023233-2222130102033223-2332103012302011-0020210112211022-1220232303021232-3032220330030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dns_ntp_config.custom_ntp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [dns_ntp_config](resources--securemesh_site_v2--reference--group-006.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- dns_ntp_config.custom_ntp

<a id="canonical-0231122333013022-1132212111001021-3313311021020221-3112100311013331-1203221122220013-0200023321212013-3333020232212123-3101320302013002"></a>

Type: `"object"`. single nested block, Optional.

NTP Servers. NTP Servers.

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
custom_ntp {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113032123102311-1001321020221113-3003321210020312-3010133031221123-1031202002221112-1301112202233202-2133112003102223-3022313311203102"></a>

### Direct properties for `dns_ntp_config.custom_ntp`

<a id="canonical-1330103310103032-3202002020101031-1010111032122310-2230033001301202-1233023202300003-3001123313300020-2331030031013022-3213002100202100"></a>

#### `dns_ntp_config.custom_ntp.ntp_servers` property

Type: `["list", "string"]`. Optional.

NTP Servers. NTP Servers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3013002232311000-0033131110202020-1213021032031033-1303213220200223-2332212011310211-1133223021130120-3333210012011000-2002210132030321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dns_ntp_config.f5_dns_default` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [dns_ntp_config](resources--securemesh_site_v2--reference--group-006.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- dns_ntp_config.f5_dns_default

<a id="canonical-1213221331011030-3021211101023201-3230303030132122-3302202321031202-1320132120320100-0322020303230101-3101101302032101-2123033331312113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for f5 DNS default.

Additional upstream details:

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
f5_dns_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013321102112123-3101002330233212-0213101211010332-1323311033231203-1022110213210023-1033323013110330-0000013023222121-0330331202122223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dns_ntp_config.f5_ntp_default` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [dns_ntp_config](resources--securemesh_site_v2--reference--group-006.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- dns_ntp_config.f5_ntp_default

<a id="canonical-2300033312132033-2122000300211013-0331323213022010-3023213210321133-2223213123111111-3111122220211212-1320002132321311-3120231102130003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for f5 ntp default.

Additional upstream details:

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
f5_ntp_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- eks_k8s

<a id="canonical-3210101303003113-2321231211232111-1200012021132302-2232223222232303-0301021301220332-2322201322313001-0323302010313111-3120122330030323"></a>

Type: `"object"`. single nested block, Optional.

Kubernetes Provider Type. Kubernetes Provider Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_anti_affinity",
    "enable_anti_affinity")}
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
  "x-ves-oneof-field-anti_affinity_choice": "[\"disable_anti_affinity\",\"enable_anti_affinity\"]"
}
```

Terraform syntax:

```terraform
eks_k8s {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002231131011001-2133302312133200-2211003220120110-0112200303001123-3103303112013202-2101020103010023-1131020112210112-0333102121331331"></a>

### Direct properties for `eks_k8s`

<a id="canonical-3021030020131220-2101123101201122-0030312122113230-0233303123321302-0333032233331010-3112121032002121-1033201220330111-0233220322022301"></a>

#### `eks_k8s.deployment_size` property

Type: `"string"`. Optional.

\[Enum: KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM|KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE\] Enum for
Kubernetes deployment size OPTIONS - KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM: Medium Medium deployment
size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most deployments. -
KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE: Large Large deployment size with higher resource.. Possible
values are \`KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM\`, \`KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE\`.
Defaults to \`KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM\`.

Additional upstream details:

Enum for Kubernetes deployment size OPTIONS

&#8203;- KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM: Medium

Medium deployment size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most
deployments. &#8203;- KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE: Large

Large deployment size with higher resource requirements (16 vCPU, 64 GB memory) for demanding
workloads requiring additional performance and capacity.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["KUBERNETES_DEPLOYMENT_SIZE_LARGE","KUBERNETES_DEPLOYMENT_SIZE_MEDIUM"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
    "KUBERNETES_DEPLOYMENT_SIZE_LARGE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
  "enum": [
    "KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
    "KUBERNETES_DEPLOYMENT_SIZE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_anti_affinity](resources--securemesh_site_v2--reference--group-006.md#canonical-2302123313002100-2233333033311201-1310100113000030-3230101310313120-2023123110023111-3310203313322012-2221021100323013-3032312333231020): complete subsection reference.

- [enable_anti_affinity](resources--securemesh_site_v2--reference--group-006.md#canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010): complete subsection reference.

<a id="canonical-1233223212313102-2112232031112231-3322120010131331-0000300213221111-0332200002020200-1010021222332203-3303131300120321-3213010031033330"></a>

<a id="canonical-1310203023023121-1131023011121211-3301300211121231-2332001102222121-3121313203133012-3003001233031122-1201002200001212-1221100212313003"></a>

#### `eks_k8s.labels` property

Type: `["map", "string"]`. Optional.

Add labels to control which Kubernetes nodes the VPM and related pods (etcd, VER, prometheus) are
deployed to. Specify label key-value pairs that match the labels on your Kubernetes nodes. This uses
Kubernetes nodeSelector to schedule pods only on nodes with matching labels.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":253,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"253\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.max_len\":\"63\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":63,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 253,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "253",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.max_len": "63",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 63,
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
    "ves.io.schema.rules.map.keys.string.max_len": "253",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "63",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "253",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "63",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010): complete subsection reference.

<a id="canonical-2302123313002100-2233333033311201-1310100113000030-3230101310313120-2023123110023111-3310203313322012-2221021100323013-3032312333231020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.disable_anti_affinity` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- eks_k8s.disable_anti_affinity

<a id="canonical-2221002233013323-1013300303313310-0213001331100022-1013213321030200-2221032231101303-0031020321332120-2232202021332210-3232313130022213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable anti affinity.

Additional upstream details:

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
disable_anti_affinity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.enable_anti_affinity` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- eks_k8s.enable_anti_affinity

<a id="canonical-2110220032230201-2223313320133330-2322130200321311-1130003031202122-2123111001031302-3022030332112010-2233021022101121-2313311201220320"></a>

Type: `"object"`. single nested block, Optional.

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
enable_anti_affinity {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300330003331000-2303302221120332-0000030222000202-1302031021303131-1231211212310123-1112120231332132-0010111002211231-0103020222103223"></a>

### Direct properties for `eks_k8s.enable_anti_affinity`

- [rules](resources--securemesh_site_v2--reference--group-006.md#canonical-0002210012021123-3003312201220103-0201333310103102-2102301312332131-0111231332122200-0321113022232322-1133333131302323-3210232203133131): complete subsection reference.

<a id="canonical-0002210012021123-3003312201220103-0201333310103102-2102301312332131-0111231332122200-0321113022232322-1133333131302323-3210232203133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.enable_anti_affinity.rules` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.enable_anti_affinity](resources--securemesh_site_v2--reference--group-006.md#canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010)
- eks_k8s.enable_anti_affinity.rules

<a id="canonical-2232011021032131-0211121302022333-3032000111100332-3033223122331020-1031220220130321-1213331313121002-2122022123203031-1223122033020313"></a>

Type: `"object"`. list nested block, Optional.

Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be
distributed across which topology domains. Example: Rule 1 - Distribute VPM pods across nodes, Rule
2 - Distribute Prometheus pods across zones.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("label_key",
    "label_value",
    "topology_keys")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121312103232000-3001311331121323-1103103012331130-3011222231311023-0121033201123320-2210023220233221-0100123211022213-2122310022103121"></a>

### Direct properties for `eks_k8s.enable_anti_affinity.rules`

<a id="canonical-3301120121212013-0202212210123121-0130011312102220-0233313230313020-3313110022032313-3003302223030333-0030032132331212-0030233121132002"></a>

#### `eks_k8s.enable_anti_affinity.rules.label_key` property

Type: `"string"`. Optional.

Specify the label key of the customer pods that CE pods should avoid being co-scheduled with.
Combined with the label value below, this identifies the target pods.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 253),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 253,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 253,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "253",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "253",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3210013103021113-0202112131113131-2303210003300232-3321000212333013-1231020203223030-2300220123021233-1012013022233031-2323231033103332"></a>

<a id="canonical-0121000312031222-1002312222310203-0113131033013232-1013210203302122-2131113032111013-0300011110111032-2323331230331003-3130232202333233"></a>

#### `eks_k8s.enable_anti_affinity.rules.label_value` property

Type: `"string"`. Optional.

Specify the label value that, together with the label key, identifies the customer pods to avoid.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0201211232020233-2123211230232103-2011011030303111-3131322111222302-1232223031321102-2233131322031330-3223232313013203-0001331120113101"></a>

<a id="canonical-3120000303100023-1231322231213123-2103010303122001-1230210222201313-0322112130130131-1100100112203312-2000100103333022-3001021002230332"></a>

#### `eks_k8s.enable_anti_affinity.rules.topology_keys` property

Type: `["list", "string"]`. Optional.

Specify one or more node label keys that define the scope of avoidance. For each topology key (e.g.,
Kubernetes.I/O/hostname), CE pods will avoid nodes whose topology value matches a node already
running a pod with the above specified label. Example: with Kubernetes.I/O/hostname, CE pods are
kept off any node running the matching customer pod.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "253",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "253",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- eks_k8s.not_managed

<a id="canonical-3130113003303110-2033010301323132-3110311003011313-1130302131021033-1223020332122222-0223032333031232-2132323111120201-1302010201330310"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1003131113220230-3313212011232303-3211321013012322-2121223030120211-3132021032320113-1233333230232103-3011113031103032-1201031231102200"></a>

### Direct properties for `eks_k8s.not_managed`

- [node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231): complete subsection reference.

<a id="canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- eks_k8s.not_managed.node_list

<a id="canonical-2122122333322033-3100113021312120-0210110100033002-1333003020021012-3020030223001111-0121002331320332-1030100303213302-1113222113030110"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0212222012323213-1111120103202002-3222032000123130-3003032021131031-3113102301131311-3203311130121201-0023221320102310-0033322112320310"></a>

### Direct properties for `eks_k8s.not_managed.node_list`

<a id="canonical-1013123311312332-0332021321113202-1022331102201200-0303333010302202-0211313031201221-1021011003120130-0221331332222012-1030310120202121"></a>

#### `eks_k8s.not_managed.node_list.hostname` property

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [interface_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210): complete subsection reference.

<a id="canonical-2333101123320133-2122103110232332-3232110303303330-0112202332322101-3301021122000131-3133030233132221-1131031112001120-3010220123210023"></a>

<a id="canonical-0313103212303232-3203133111211001-1031222331101222-3230102130033222-2123012221313201-2130230112310201-1010303332012021-3221213223032022"></a>

#### `eks_k8s.not_managed.node_list.public_ip` property

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1013211111013203-3031113031113313-1332002303201120-2331233311113020-3322121013003110-0301023020032022-0021001021312301-0222230301020222"></a>

<a id="canonical-0321013300232110-2101333030110310-2233310003310323-0003001130022233-2323101113331213-0012132010202033-1322210322201110-2330131003233200"></a>

#### `eks_k8s.not_managed.node_list.type` property

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Control","Worker"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-006.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-006.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- eks_k8s.not_managed.node_list.interface_list

<a id="canonical-1213312202021031-3002103122300032-3303202132111103-0323331330231311-1103123011021111-0012232113202222-2302200112122320-1313232012013121"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3313100320203012-0230002221000033-3203132031302113-3330300302101103-2320213013303000-3312302001323303-0132001233212033-0020102112213320"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331): complete subsection reference.

<a id="canonical-2231213231231020-3201113010022121-1103331331121023-2332030200332231-0113020332210120-3131003133320020-1022231113322310-1022003133132322"></a>

<a id="canonical-1210001101003012-0111113010300031-2221110010032122-3323103130112002-0331321311000330-2032030011120320-1011030231102333-1112322131003310"></a>

#### `eks_k8s.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-007.md#canonical-3012211120012130-2120220310021002-0011001021013213-1201330020123013-2321130021032122-1223300321221131-2310211131333311-0310233021221020): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-1332210212311330-2022022112011131-1010231320331201-2022030010222231-3332011213032033-0302230033003033-0112221032213112-1223000100220010): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222): complete subsection reference.

<a id="canonical-0330221100330133-2330201002202021-1020021311111200-3013201210033003-3221121101123030-3023011312210233-3321020100213102-1313032302232122"></a>

<a id="canonical-3222111302021222-3100010330002133-0112021311112130-1313031012331113-3222330123131001-0333212031313023-1000122110132310-3002220030101203"></a>

#### `eks_k8s.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0101321101302313-3311010012322323-3222232321002222-1310222112033112-0323302112332202-3132000121211231-2000232130211012-0331332021311031"></a>

<a id="canonical-1310010003111200-3032012011203112-2032302030331033-3003132221110001-0203230322012101-0111132123122201-2312322210301031-0302313203003010"></a>

#### `eks_k8s.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1123103321332322-0320210301132223-3332113001103202-2012111310113013-0013111030113202-2331210011332221-1132330020131232-1300202312220202"></a>

<a id="canonical-3231112302022120-2231012011323311-0102332111101222-3032230133103110-1122331111212113-3110110300233311-0012300133211203-3023122110230313"></a>

#### `eks_k8s.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [monitor](resources--securemesh_site_v2--reference--group-007.md#canonical-0011102111120011-1022033123221032-1331321200332203-2301111122122233-3303333213003203-2021123332202300-1022321002112333-0023222221123022): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-007.md#canonical-3203001200120111-0313132010231330-3001221200201310-3122003203211002-0032332311222332-1021320120202100-3231020003010120-3230012212102130): complete subsection reference.

<a id="canonical-0301030223013230-0121223030131313-3123231100031000-0222112012031012-3220211322111231-2233210300033111-2012001321022020-3102132321102022"></a>

<a id="canonical-0012123331220211-1331310002203123-2122333000300333-3232000100131221-0210300311230232-1023321010210231-3231231102121110-0201222222232312"></a>

#### `eks_k8s.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-1320213111202200-3223222232333220-3310033133012301-3200103100012113-1130332200033123-3110002222310212-2232312022122322-1221200223121303"></a>

<a id="canonical-2333203003110132-2122030232130030-0122022112023230-0011020000203302-2110122200011313-1023221123110003-0023331310230013-3223102122023200"></a>

#### `eks_k8s.not_managed.node_list.interface_list.name` property

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [network_option](resources--securemesh_site_v2--reference--group-007.md#canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2301031122023323-2312030233133232-0003013033030323-3223231102220222-3121003131102022-1010213001332131-0132020321212203-3121010111120023): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-1130210113011010-0201323203331131-2111202121232233-0002232003002002-1113221321222220-3032221120013120-1301312312021220-3211333120012012): complete subsection reference.

<a id="canonical-0231332122113033-0023022123003230-1130213133212122-1020213220303111-2030323333210223-1312023322312323-3133133330012323-3133210130300330"></a>
