---
page_title: "xcsh_k8s_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_cluster reference."
---

# xcsh_k8s_cluster reference

<a id="canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- Property reference

<a id="canonical-1030112223322233-0231211231003320-3020002220320010-1332331133131302-1130312231102132-0031310312220311-0233312221301100-2031213120220301"></a>

### Direct properties for `xcsh_k8s_cluster`

<a id="canonical-0033211102013120-3021022321300103-2223032023013220-0101331303311100-0023121100223210-2300131200330333-0331113122313013-1220101210113311"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
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
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [cluster_scoped_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-3231202212112233-0112112212220131-2001200010200310-1222310220230102-3120032033032230-0112121222011032-2331302013101331-2103103120021321): complete subsection reference.

- [cluster_scoped_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-1332331123012022-0333331301220120-3313321012131011-3322323330321232-0310123231231331-1311023132132232-0201003010213010-1013302023201311): complete subsection reference.

- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030): complete subsection reference.

<a id="canonical-3120131310021112-1303120032223122-0032011103223231-0111122310113232-1321300323213223-0121201222110022-2312012331202320-0200013320321122"></a>

<a id="canonical-3312300231132302-0110000003132321-2110233123100320-3231233311333133-3031023132123021-1012121323010133-1030303020131233-2011333320021211"></a>

#### `description` property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-0230300110310032-0312122121123010-2020310022210323-1220231302232220-0131232220200132-1002212200312233-2012000223230320-1132002032002113"></a>

<a id="canonical-2231111000303202-1131331110302313-1221212300301222-3332223313030123-2310201320020201-1331233120221131-2033032333111001-0233321332210102"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

- [global_access_enable](resources--k8s_cluster--reference--group-001.md#canonical-1110033101330033-3012032322223221-3012123023320130-1102100010210112-1112321122232211-0033011130032123-1102210111203133-1110010101001011): complete subsection reference.

<a id="canonical-1201333221120333-2031032030310133-0311311302232300-1312313332013133-2110222310010332-2331311323303001-0333100230202100-1310033333100000"></a>

<a id="canonical-3213101210121212-0000110302222222-3212310100011001-1001231011010211-2220122300323231-1310102331211130-1300302010131223-2302311101011211"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [insecure_registry_list](resources--k8s_cluster--reference--group-001.md#canonical-2133012030210030-0220221330213223-2132132333211220-3120030220103203-1313003032301101-1300133311312022-1321021211021222-3230112201130332): complete subsection reference.

<a id="canonical-2312132012012211-3213032001021120-2022323323122322-1111313332111330-0033020202311213-3313112132133031-3020010132322321-1110022210112303"></a>

<a id="canonical-3003122020022312-1320032030311200-3330013121102010-1003131121200310-1320322131100232-0031011303310330-0211311310103310-3003023221322212"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

- [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-2330113202310232-1313112231102321-0100021023331021-3030000310021310-0120111130300032-1301121303202001-0330002111131003-3333210313223011): complete subsection reference.

<a id="canonical-2213320103200211-2001212001301023-2132321103110310-3211131033322230-3023312022322201-2301230032102033-2032002200330311-1022021132110131"></a>

<a id="canonical-1331002200113111-3223210002122123-3312232113331330-2300331000112101-3112230232301312-3122303103200332-3123332130232332-3001103220321113"></a>

#### `name` property

Type: `"string"`. Required.

Name of the K8S Cluster. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0231000200020332-0110303221012123-1310102302310132-3122102300321300-3020213222032213-0102211200020300-3012330100111121-1022310202110302"></a>

<a id="canonical-0120230032003223-1302123302132132-3203023210313120-3201300210030312-0203301030301122-3030200201320011-2303110003303001-0222123101100130"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the K8S Cluster. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
EnumExtractionComplete: false
EnumValidators: [{"version":1,"validator":"OneOf","values":["system"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [no_cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-2210131033311120-1312231122121133-3230223112010200-2121132220123123-1121312322203212-2333231201000112-3202301221212121-1221201220201322): complete subsection reference.

- [no_global_access](resources--k8s_cluster--reference--group-001.md#canonical-3102111002233003-1113331000033312-3121001211313133-2221120110302331-1321010313323021-0312101222012023-3103201132033230-1022003113010030): complete subsection reference.

- [no_insecure_registries](resources--k8s_cluster--reference--group-001.md#canonical-1032123322021013-3102121121302220-1301312300210220-0031131220103302-1310012310132313-3330231103033230-2313301123001103-3333332000300300): complete subsection reference.

- [no_local_access](resources--k8s_cluster--reference--group-001.md#canonical-2000131320001322-3231301312222232-0321232331030302-1323131212032120-1021233030121121-0012033021120311-3332231013132102-3210020112130130): complete subsection reference.

- [timeouts](resources--k8s_cluster--reference--group-001.md#canonical-3222111132220221-2003132131331103-2132032010321023-0230211030303331-1333110222213201-2121103131130022-3021231323211331-2010330002322221): complete subsection reference.

- [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-0211302322000100-0203322011230121-0312133132101320-1031210001131331-1031230221102200-3221230333033330-3133111123333321-0133230211020211): complete subsection reference.

- [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-2033312333111003-3123320312002221-3221201213321321-0213023322032121-3232230123201131-0222231311213030-3330211220123210-0223222330320322): complete subsection reference.

- [use_custom_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-2312030321230203-3200103233112030-1101003101310211-1303122131333011-3312132231111332-1031233123320221-1203002202101313-2123320021313222): complete subsection reference.

- [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-3131020221100132-1003100033303012-3130330301313330-1033003233300201-1331103000203002-0033021330203230-3013312030313301-0011311101333322): complete subsection reference.

- [use_default_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-3023013200022221-2131212202021103-2232231332300220-1010102031222320-2310231313121132-0120123013030132-1233103103321123-1213221201311211): complete subsection reference.

- [use_default_cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-0213233222332312-2010100313000111-1130013031231223-0203332222133032-1321023231023111-0132310322222130-0213113132211321-0130123032022033): complete subsection reference.

- [use_default_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-2302100303223331-0223211122120012-2020001102301202-0333121332233033-0310133202333012-0111223110010210-1212202201010123-3322101332123102): complete subsection reference.

- [use_default_psp](resources--k8s_cluster--reference--group-001.md#canonical-0301302001120301-3120213203003133-3013101332200030-3303203223000211-1100230021000103-3230002012102133-0311030302011303-2322000302222001): complete subsection reference.

- [vk8s_namespace_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-2210013331002303-1110100111123032-0010132331001203-0120302213302320-1110231110012230-3113330311303112-3232113213313233-1023013222120221): complete subsection reference.

- [vk8s_namespace_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-2011031013012102-2201320113333031-2201203010203212-0000332201221301-1301311203111302-2022113130022030-1222210102113232-3302123213330233): complete subsection reference.

<a id="canonical-0331111021221011-0023021230021033-0110320022303002-2001102320331012-1211010003100121-2120030230231122-1110202211321121-0000111333103001"></a>

### All schema paths for `xcsh_k8s_cluster`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_cluster--reference--group-001.md#canonical-0033211102013120-3021022321300103-2223032023013220-0101331303311100-0023121100223210-2300131200330333-0331113122313013-1220101210113311) |
| `cluster_scoped_access_deny` | [cluster_scoped_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-0133203001001103-1222111320011300-2320203211222222-1012012323002131-2231200102003021-1303310112100112-0312000312100331-1212102123200123) |
| `cluster_scoped_access_permit` | [cluster_scoped_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-2223313031321222-0003211003311222-0311212122110333-3321213022321123-2030032201032130-1231013122032132-3201103232300301-3201212112323103) |
| `cluster_wide_app_list` | [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-2223302213100033-1212333122233021-1231233202330023-0301323021130332-2300331000220201-0131313302332111-0332220320212211-1122021301101210) |
| `cluster_wide_app_list.cluster_wide_apps` | [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0011230011103013-2023233223213110-2102220303202021-3311011231222023-2230223121101101-2130201322203210-1312113133103010-3222123101031123) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd` | [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-3233331310201300-2211321121322222-2233330033111013-3333130202223301-3012321201310313-3322023031223221-0232320333001023-0322033033331323) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-2000303131032300-0322120023310032-3311321002321003-3311130100302002-3210102223103030-3320131000022132-0323021211113331-2213231031121011) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port](resources--k8s_cluster--reference--group-001.md#canonical-3212030222033221-3102211320311303-3003312022220001-3211212211333231-3132211233011022-2111202000100110-2303211201332003-1033112222213030) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-3101102330132132-0312322220000300-1120233103232201-0303111212033133-3000003030313223-0201223023330120-3103122101111310-0131201132202201) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--reference--group-001.md#canonical-0133331113320211-0213102210210123-1332332012031001-3103033010100310-2220000332111022-2003110031301321-1223112233202220-2210220030123301) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-3121220213113210-3310203010002320-0031000333213101-3102323032003311-1311323232322111-1102321111131223-3013211023121023-3202210013112300) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider](resources--k8s_cluster--reference--group-001.md#canonical-3213223131223200-0031021302212031-3101133131202121-0002122013112230-1203333300111111-2112023231202202-1232131010331202-1323110320120100) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location](resources--k8s_cluster--reference--group-001.md#canonical-3012213032110031-3220022231133210-3230311000320323-0012030013023120-3110323132233202-3012210100233001-1322333321212332-2202223202302201) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider](resources--k8s_cluster--reference--group-001.md#canonical-1031112113330233-2122101100021323-0121030302233110-3301313001121030-3102230012222003-0233301001001320-3110023303032001-3000212232030021) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-1133210222130122-3120330223002320-0230230202233310-1230023023233011-2301121021131122-3131003001210102-2012021031021033-1122122031212002) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref](resources--k8s_cluster--reference--group-001.md#canonical-1032002003100032-3030232203011033-1321321232120200-0303112220012320-2312032220110200-3011110213101230-2222131130322220-0102320011301120) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url](resources--k8s_cluster--reference--group-001.md#canonical-0110333030330122-1021322010022221-2103100000232022-3232032332020002-1201133322030202-0202123111033020-0103333113122020-1202312213020301) |
| `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port` | [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port](resources--k8s_cluster--reference--group-001.md#canonical-3203012311132301-0212030033103232-0332012202232023-1133000301131110-0211233033211123-0121003000010310-1013300312331302-1332231212002002) |
| `cluster_wide_app_list.cluster_wide_apps.dashboard` | [cluster_wide_app_list.cluster_wide_apps.dashboard](resources--k8s_cluster--reference--group-001.md#canonical-2232032312201122-3213220312121003-1023011133202202-2230003110200201-3303231003223133-3023220113222003-3002330213013332-3220333100320131) |
| `cluster_wide_app_list.cluster_wide_apps.metrics_server` | [cluster_wide_app_list.cluster_wide_apps.metrics_server](resources--k8s_cluster--reference--group-001.md#canonical-1000230123303313-0313331023011132-3003312220203332-1120231202221030-0103113133323001-0322333222102321-1123212003122121-2021100021303013) |
| `cluster_wide_app_list.cluster_wide_apps.prometheus` | [cluster_wide_app_list.cluster_wide_apps.prometheus](resources--k8s_cluster--reference--group-001.md#canonical-1230310330313303-3230203220201201-0120312220113120-1313223300331131-3221320130000001-0002032210110122-0203320211211132-0223213002001110) |
| `description` | [description](resources--k8s_cluster--reference--group-001.md#canonical-3120131310021112-1303120032223122-0032011103223231-0111122310113232-1321300323213223-0121201222110022-2312012331202320-0200013320321122) |
| `disable` | [disable](resources--k8s_cluster--reference--group-001.md#canonical-0230300110310032-0312122121123010-2020310022210323-1220231302232220-0131232220200132-1002212200312233-2012000223230320-1132002032002113) |
| `global_access_enable` | [global_access_enable](resources--k8s_cluster--reference--group-001.md#canonical-0213110233322333-0223333313210033-2200331030213133-3221131213003323-1322122010031300-0120010011230331-1122030002300300-3132031023200210) |
| `id` | [ID](resources--k8s_cluster--reference--group-001.md#canonical-1201333221120333-2031032030310133-0311311302232300-1312313332013133-2110222310010332-2331311323303001-0333100230202100-1310033333100000) |
| `insecure_registry_list` | [insecure_registry_list](resources--k8s_cluster--reference--group-001.md#canonical-2100022003302111-1011113003200323-0001000100202102-3201032300101100-0133302230230000-0301133020103022-0121331223222130-2331233313210210) |
| `insecure_registry_list.insecure_registries` | [insecure_registry_list.insecure_registries](resources--k8s_cluster--reference--group-001.md#canonical-0012112211200110-1210123133001133-0113030332200132-0312301210200120-1100233121233123-3130320022223321-2232232212101130-0333233233301121) |
| `labels` | [labels](resources--k8s_cluster--reference--group-001.md#canonical-2312132012012211-3213032001021120-2022323323122322-1111313332111330-0033020202311213-3313112132133031-3020010132322321-1110022210112303) |
| `local_access_config` | [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-2103201032303300-1320122020201231-0303123212022321-0132331230310013-2223322213330222-2132131001211030-0103132110110301-0311100111033223) |
| `local_access_config.default_port` | [local_access_config.default_port](resources--k8s_cluster--reference--group-001.md#canonical-1011003110233320-2210113003110302-1333111133231213-0203120302002012-1011022213120223-1301211232120322-1332210223321120-2333213101123231) |
| `local_access_config.local_domain` | [local_access_config.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-3300111003312031-0131021221003110-2001203120302010-3211130331310032-1131210032102111-1320001020320322-3220232020322100-3102312112010222) |
| `local_access_config.port` | [local_access_config.port](resources--k8s_cluster--reference--group-001.md#canonical-3010202023123233-1201310032301120-0030303113022002-2002330211232020-2221010320311103-1222232331333320-1111020021230031-2310013013011232) |
| `name` | [name](resources--k8s_cluster--reference--group-001.md#canonical-2213320103200211-2001212001301023-2132321103110310-3211131033322230-3023312022322201-2301230032102033-2032002200330311-1022021132110131) |
| `namespace` | [namespace](resources--k8s_cluster--reference--group-001.md#canonical-0231000200020332-0110303221012123-1310102302310132-3122102300321300-3020213222032213-0102211200020300-3012330100111121-1022310202110302) |
| `no_cluster_wide_apps` | [no_cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-2301311322321010-0011123230311100-2331130103233323-2121301212012303-3211131023133301-0011123330122231-1220100310213211-3302233100001212) |
| `no_global_access` | [no_global_access](resources--k8s_cluster--reference--group-001.md#canonical-3313313210202332-2111101111001300-0111010313003003-3122222132131112-3120213213332121-3113331200233203-0112110330033330-3110332002001012) |
| `no_insecure_registries` | [no_insecure_registries](resources--k8s_cluster--reference--group-001.md#canonical-3321121321223302-1210011303220123-3032020300021012-1220022120321133-2222201132113201-0310002312231202-1303320031013001-1211100113132101) |
| `no_local_access` | [no_local_access](resources--k8s_cluster--reference--group-001.md#canonical-3110122132203203-0211212131133312-3331101212003103-0320210130101012-2211012100230112-2130110102310220-2303213102030221-2101020110023220) |
| `timeouts` | [timeouts](resources--k8s_cluster--reference--group-001.md#canonical-3210331111102101-2212302231321020-0023102223223201-0133032032332121-1331301302010010-1111232011232110-1132130313231020-2132230232101330) |
| `timeouts.create` | [timeouts.create](resources--k8s_cluster--reference--group-001.md#canonical-0120300022120030-0023201020020203-1120302213321233-1321220200102032-3123333302110201-1312031023131310-2102033110320202-0210212032231100) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_cluster--reference--group-001.md#canonical-1322331131212113-0132003002013200-3013011111010101-1112132213122113-2022120220200220-2122132232300133-1030211233003013-0332231002210213) |
| `timeouts.read` | [timeouts.read](resources--k8s_cluster--reference--group-001.md#canonical-1312130332312022-1232131323031203-0212230123223000-2320230333321320-2230301211333103-0000213230103110-3230230121312123-0110132300003012) |
| `timeouts.update` | [timeouts.update](resources--k8s_cluster--reference--group-001.md#canonical-2003122221200031-1030011311220003-3202220313201332-1312211312330103-0023311311330120-0111130122233130-3010122212002023-2033202310300231) |
| `use_custom_cluster_role_bindings` | [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-1021121302311233-3033003300301332-2021111222212201-3033121221133222-3110101231113213-1133002131200303-2030303112330200-2000330133330122) |
| `use_custom_cluster_role_bindings.cluster_role_bindings` | [use_custom_cluster_role_bindings.cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-1311103030322321-2301132231011212-3122022103002011-1100302033021120-2323211213022101-1020112323022111-1112313312312111-3323002322133231) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.name` | [use_custom_cluster_role_bindings.cluster_role_bindings.name](resources--k8s_cluster--reference--group-001.md#canonical-0203102312330303-3333230223223313-0113233031102101-2333223001303223-1232022113012032-0330202012311120-0032222321231032-1000131032133222) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.namespace` | [use_custom_cluster_role_bindings.cluster_role_bindings.namespace](resources--k8s_cluster--reference--group-001.md#canonical-3121101103122021-3311132023113021-3123221333233232-1030022331332202-1310312030033012-1122211031100033-0303030010003011-2212332300313221) |
| `use_custom_cluster_role_bindings.cluster_role_bindings.tenant` | [use_custom_cluster_role_bindings.cluster_role_bindings.tenant](resources--k8s_cluster--reference--group-001.md#canonical-0301203123102130-0231003020330033-1000013311211323-0100123031232333-0302120220112303-1202322131320230-0231300000123103-2012303333320333) |
| `use_custom_cluster_role_list` | [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-0322303213323223-3200020033032230-3000200200031113-2021020211322120-3313230231121231-1300231021300120-3311030230302222-3302201212112032) |
| `use_custom_cluster_role_list.cluster_roles` | [use_custom_cluster_role_list.cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-0222330201030200-1130023021002111-0210311020013310-3013113233012111-1303003311202112-0113010332322022-2031132232130211-3322022311311230) |
| `use_custom_cluster_role_list.cluster_roles.name` | [use_custom_cluster_role_list.cluster_roles.name](resources--k8s_cluster--reference--group-001.md#canonical-1031302202112211-1211012023213132-0022233103321222-3023312011020101-1133322113200220-1201211101210100-3011310012021103-3233232103122301) |
| `use_custom_cluster_role_list.cluster_roles.namespace` | [use_custom_cluster_role_list.cluster_roles.namespace](resources--k8s_cluster--reference--group-001.md#canonical-0210223200023213-2111332212101103-1231132121322311-1122302213121100-3220223132012131-1123300031301112-3330320031022122-1023010033013123) |
| `use_custom_cluster_role_list.cluster_roles.tenant` | [use_custom_cluster_role_list.cluster_roles.tenant](resources--k8s_cluster--reference--group-001.md#canonical-3102233333010133-2232331001031333-0022032122203123-1333230113122212-1312021110031031-0130212200100303-3112213332330033-2203210313122113) |
| `use_custom_pod_security_admission` | [use_custom_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-3300233020112301-2311123202333232-2031121320200311-2313313022022230-0212210211003011-3211022120222213-1320213322100212-3231131212202103) |
| `use_custom_pod_security_admission.name` | [use_custom_pod_security_admission.name](resources--k8s_cluster--reference--group-001.md#canonical-1313032303103201-3322302221232223-0333201122010130-0003123003321030-2201233220021231-2320322210213120-1211022322131323-2021021021012110) |
| `use_custom_pod_security_admission.namespace` | [use_custom_pod_security_admission.namespace](resources--k8s_cluster--reference--group-001.md#canonical-2000101231132133-0010320312323122-2022111233313212-3210132223330330-2020111222221023-1021332033002300-2303130230201213-2002332133303303) |
| `use_custom_pod_security_admission.tenant` | [use_custom_pod_security_admission.tenant](resources--k8s_cluster--reference--group-001.md#canonical-3331032112211123-0320012323301322-3100112112103010-1002111131322303-2101030101023120-1232221002023303-2302211113112311-0313121310223312) |
| `use_custom_psp_list` | [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-1131122010200101-0301221002021212-2031120220103122-1220101122031233-3030210323122033-1032313223020321-0131221300120122-3103332022210303) |
| `use_custom_psp_list.pod_security_policies` | [use_custom_psp_list.pod_security_policies](resources--k8s_cluster--reference--group-001.md#canonical-2322203130113112-1101133031321302-0301222330212123-2113003102221210-0211312302311300-1202203213200320-1311222100212132-2333232111330320) |
| `use_custom_psp_list.pod_security_policies.name` | [use_custom_psp_list.pod_security_policies.name](resources--k8s_cluster--reference--group-001.md#canonical-1101032010301010-2213131112302211-2312313211231221-0302032301121312-3120222221120300-0110011221233330-1202220222131223-2312213323303312) |
| `use_custom_psp_list.pod_security_policies.namespace` | [use_custom_psp_list.pod_security_policies.namespace](resources--k8s_cluster--reference--group-001.md#canonical-3112211333030212-2332300023013002-0120132110230233-3011312200113133-3320111231233332-3020132030222303-1013223310311210-2100122030230313) |
| `use_custom_psp_list.pod_security_policies.tenant` | [use_custom_psp_list.pod_security_policies.tenant](resources--k8s_cluster--reference--group-001.md#canonical-2311132231222210-3201120210023112-3211022213033002-0112312111101301-1110000011100300-2002120223033320-1123132201132223-0031201312221120) |
| `use_default_cluster_role_bindings` | [use_default_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-2333331010123001-0300022310113220-2120001233210003-0213233233310313-1133030023130213-1311312002033201-1020133232011102-3111311012213101) |
| `use_default_cluster_roles` | [use_default_cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-3321330210122120-1033210002331112-0300110101203330-3230233323313033-2032012001212121-1203031310100031-2002120323003110-2332013121311321) |
| `use_default_pod_security_admission` | [use_default_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-0100202212033302-2322133002120211-0231000210210323-0101012210103120-1320101031313213-2132222030100232-2103020200102333-1211300022221301) |
| `use_default_psp` | [use_default_psp](resources--k8s_cluster--reference--group-001.md#canonical-2330110111230123-1223231130133333-2011031001010130-3231111012320030-0023322012102020-1003002331130312-1130001312131101-2331023211112012) |
| `vk8s_namespace_access_deny` | [vk8s_namespace_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-1230112122021332-1321330221320110-3030121021021303-3032313030210320-3231101112320131-0302120333312322-3012310121323302-1322021020212110) |
| `vk8s_namespace_access_permit` | [vk8s_namespace_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-1210200220020311-0221232223021310-2011020012103310-1111201321220202-3010002102211322-0100312231011033-1322210023021220-0131212330010113) |

<a id="canonical-3231202212112233-0112112212220131-2001200010200310-1222310220230102-3120032033032230-0112121222011032-2331302013101331-2103103120021321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_scoped_access_deny` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- cluster_scoped_access_deny

<a id="canonical-0133203001001103-1222111320011300-2320203211222222-1012012323002131-2231200102003021-1303310112100112-0312000312100331-1212102123200123"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: cluster\_scoped\_access\_deny, cluster\_scoped\_access\_permit\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [cluster_scoped_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-0133203001001103-1222111320011300-2320203211222222-1012012323002131-2231200102003021-1303310112100112-0312000312100331-1212102123200123)
- [cluster_scoped_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-2223313031321222-0003211003311222-0311212122110333-3321213022321123-2030032201032130-1231013122032132-3201103232300301-3201212112323103)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cluster_scoped_access_deny = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332331123012022-0333331301220120-3313321012131011-3322323330321232-0310123231231331-1311023132132232-0201003010213010-1013302023201311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_scoped_access_permit` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- cluster_scoped_access_permit

<a id="canonical-2223313031321222-0003211003311222-0311212122110333-3321213022321123-2030032201032130-1231013122032132-3201103232300301-3201212112323103"></a>

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
cluster_scoped_access_permit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- cluster_wide_app_list

<a id="canonical-2223302213100033-1212333122233021-1231233202330023-0301323021130332-2300331000220201-0131313302332111-0332220320212211-1122021301101210"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: cluster\_wide\_app\_list, no\_cluster\_wide\_apps; Default: no\_cluster\_wide\_apps\]
Cluster Wide Application List. List of cluster wide applications.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_wide_apps")}
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

- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-2223302213100033-1212333122233021-1231233202330023-0301323021130332-2300331000220201-0131313302332111-0332220320212211-1122021301101210)
- [no_cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-2301311322321010-0011123230311100-2331130103233323-2121301212012303-3211131023133301-0011123330122231-1220100310213211-3302233100001212)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cluster_wide_app_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203003311321110-2203303222032112-1023300023101300-0032221030131220-0110031303221310-1232331012232322-2231021323002320-0222302020013231"></a>

### Direct properties for `cluster_wide_app_list`

- [cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020): complete subsection reference.

<a id="canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- cluster_wide_app_list.cluster_wide_apps

<a id="canonical-0011230011103013-2023233223213110-2102220303202021-3311011231222023-2230223121101101-2130201322203210-1312113133103010-3222123101031123"></a>

Type: `"object"`. list nested block, Optional.

Cluster Wide Application List. List of cluster wide applications.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("argo_cd",
    "dashboard"),
  validators.ConflictingListObjectAttributes("argo_cd",
    "metrics_server"),
  validators.ConflictingListObjectAttributes("argo_cd",
    "prometheus"),
  validators.ConflictingListObjectAttributes("dashboard",
    "metrics_server"),
  validators.ConflictingListObjectAttributes("dashboard",
    "prometheus"),
  validators.ConflictingListObjectAttributes("metrics_server",
    "prometheus")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
cluster_wide_apps {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101012303300230-3013213231233332-3120232101230033-1303022012011100-2020122010202203-2330131001212202-0210021213312031-3300030003223102"></a>

### Direct properties for `cluster_wide_app_list.cluster_wide_apps`

- [argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-0332100103210130-3303021300221232-1002012211332320-2011210220030111-0300222103122203-3300233010303203-0233111303210230-3033131021312230): complete subsection reference.

- [dashboard](resources--k8s_cluster--reference--group-001.md#canonical-3111130013310203-2021310113320211-1112100020113323-3123101301130332-1301211221103223-0123110011020110-0002210120232023-3010031222033000): complete subsection reference.

- [metrics_server](resources--k8s_cluster--reference--group-001.md#canonical-0031011133333033-1333111103030230-3032320110232222-3122022012122220-1202203013332031-2033031112321320-1311302221323121-2220010113000321): complete subsection reference.

- [prometheus](resources--k8s_cluster--reference--group-001.md#canonical-2102233013113032-1120130222210211-2313102122223131-0212311203003303-2121032202130330-3000012030231030-2300301033202210-3023131302301013): complete subsection reference.

<a id="canonical-0332100103210130-3303021300221232-1002012211332320-2011210220030111-0300222103122203-3300233010303203-0233111303210230-3033131021312230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps.argo_cd` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020)
- cluster_wide_app_list.cluster_wide_apps.argo_cd

<a id="canonical-3233331310201300-2211321121322222-2233330033111013-3333130202223301-3012321201310313-3322023031223221-0232320333001023-0322033033331323"></a>

Type: `"object"`. single nested block, Optional.

Description Parameters for Argo Continuous Deployment(CD) application.

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
argo_cd {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113223322020031-0112112000001200-1222102030020233-3322110010002230-3031021131303333-3020231233010021-2011213222113310-1003020322122233"></a>

### Direct properties for `cluster_wide_app_list.cluster_wide_apps.argo_cd`

- [local_domain](resources--k8s_cluster--reference--group-001.md#canonical-2110100101020000-2202001310313322-3101301131020012-3103330021111120-0301122101212132-0323230123013221-3212230112312322-2330320300130301): complete subsection reference.

<a id="canonical-2110100101020000-2202001310313322-3101301131020012-3103330021111120-0301122101212132-0323230123013221-3212230112312322-2330320300130301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-0332100103210130-3303021300221232-1002012211332320-2011210220030111-0300222103122203-3300233010303203-0233111303210230-3033131021312230)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain

<a id="canonical-2000303131032300-0322120023310032-3311321002321003-3311130100302002-3210102223103030-3320131000022132-0323021211113331-2213231031121011"></a>

Type: `"object"`. single nested block, Optional.

Parameters required to enable local access.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("local_domain"),
  validators.ConflictingObjectAttributes("default_port",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

Terraform syntax:

```terraform
local_domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331303323103213-3303023320120232-1222032202002331-2231202110212233-1312001233211211-3110232301333203-2021232223312003-1302011203030012"></a>

### Direct properties for `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain`

- [default_port](resources--k8s_cluster--reference--group-001.md#canonical-1112303202332110-2030122010303220-1233133311311100-1122331100301131-0331100121330020-3312233033020102-2302202210313100-0102321223200211): complete subsection reference.

<a id="canonical-3101102330132132-0312322220000300-1120233103232201-0303111212033133-3000003030313223-0201223023330120-3103122101111310-0131201132202201"></a>

<a id="canonical-2203110331313112-3322002133231001-3023212322120221-0313333333313301-2030023230033033-1033003023321230-3100130132200220-0302332030100323"></a>

#### `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.local_domain` property

Type: `"string"`. Optional.

ArgoCD will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [password](resources--k8s_cluster--reference--group-001.md#canonical-3221232020303001-3003231303331233-1132210113332003-1211211323113310-3203312110221013-1332122020311003-0202033313310012-2313111303012003): complete subsection reference.

<a id="canonical-3203012311132301-0212030033103232-0332012202232023-1133000301131110-0211233033211123-0121003000010310-1013300312331302-1332231212002002"></a>

<a id="canonical-1112013031113233-2000002001103203-2012330000101101-1332201200100110-0303133222110031-1231003210000003-0320330033013212-2202202131330303"></a>

#### `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.port` property

Type: `"number"`. Optional.

Exclusive with \[default\_port\] Use custom ArgoCD port. Available port range is less than 65000
except reserved ports.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  }
}
```

<a id="canonical-1112303202332110-2030122010303220-1233133311311100-1122331100301131-0331100121330020-3312233033020102-2302202210313100-0102321223200211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-0332100103210130-3303021300221232-1002012211332320-2011210220030111-0300222103122203-3300233010303203-0233111303210230-3033131021312230)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-2110100101020000-2202001310313322-3101301131020012-3103330021111120-0301122101212132-0323230123013221-3212230112312322-2330320300130301)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.default_port

<a id="canonical-3212030222033221-3102211320311303-3003312022220001-3211212211333231-3132211233011022-2111202000100110-2303211201332003-1033112222213030"></a>

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
default_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221232020303001-3003231303331233-1132210113332003-1211211323113310-3203312110221013-1332122020311003-0202033313310012-2313111303012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-0332100103210130-3303021300221232-1002012211332320-2011210220030111-0300222103122203-3300233010303203-0233111303210230-3033131021312230)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-2110100101020000-2202001310313322-3101301131020012-3103330021111120-0301122101212132-0323230123013221-3212230112312322-2330320300130301)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password

<a id="canonical-0133331113320211-0213102210210123-1332332012031001-3103033010100310-2220000332111022-2003110031301321-1223112233202220-2210220030123301"></a>

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

<a id="canonical-2213001310030232-3211321011300213-2201203130010331-0020312223323001-2020110220103110-3221303120131100-2203130010200302-1323200000021323"></a>

### Direct properties for `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password`

- [blindfold_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-0331222323220332-2131310100120123-3331130321332300-1301320212112010-0132012221020212-0031011023321232-0012331200013222-3022111220300131): complete subsection reference.

- [clear_secret_info](resources--k8s_cluster--reference--group-001.md#canonical-1023210333221303-0012111313301202-2310003212030110-3101223212233333-3220132332033220-0233133323300100-3333333311331322-1011010020310332): complete subsection reference.

<a id="canonical-0331222323220332-2131310100120123-3331130321332300-1301320212112010-0132012221020212-0031011023321232-0012331200013222-3022111220300131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-0332100103210130-3303021300221232-1002012211332320-2011210220030111-0300222103122203-3300233010303203-0233111303210230-3033131021312230)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-2110100101020000-2202001310313322-3101301131020012-3103330021111120-0301122101212132-0323230123013221-3212230112312322-2330320300130301)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--reference--group-001.md#canonical-3221232020303001-3003231303331233-1132210113332003-1211211323113310-3203312110221013-1332122020311003-0202033313310012-2313111303012003)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info

<a id="canonical-3121220213113210-3310203010002320-0031000333213101-3102323032003311-1311323232322111-1102321111131223-3013211023121023-3202210013112300"></a>

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

<a id="canonical-3010102111223002-2032301213212132-0220020103131023-1330011133201011-2111332022331211-3111201332032220-3302323100213102-2312132231100330"></a>

### Direct properties for `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info`

<a id="canonical-3213223131223200-0031021302212031-3101133131202121-0002122013112230-1203333300111111-2112023231202202-1232131010331202-1323110320120100"></a>

#### `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3012213032110031-3220022231133210-3230311000320323-0012030013023120-3110323132233202-3012210100233001-1322333321212332-2202223202302201"></a>

<a id="canonical-1020110223100120-1331333221021220-2312103113213001-3111333203133111-3320033001132032-1223112111103131-3213132200011230-3000032300133230"></a>

#### `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1031112113330233-2122101100021323-0121030302233110-3301313001121030-3102230012222003-0233301001001320-3110023303032001-3000212232030021"></a>

<a id="canonical-1232233013012132-1330221030030323-2231121312002210-1323113023310221-0123222330013330-2030202023330022-0310033132103112-2332200220220220"></a>

#### `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1023210333221303-0012111313301202-2310003212030110-3101223212233333-3220132332033220-0233133323300100-3333333311331322-1011010020310332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd](resources--k8s_cluster--reference--group-001.md#canonical-0332100103210130-3303021300221232-1002012211332320-2011210220030111-0300222103122203-3300233010303203-0233111303210230-3033131021312230)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain](resources--k8s_cluster--reference--group-001.md#canonical-2110100101020000-2202001310313322-3101301131020012-3103330021111120-0301122101212132-0323230123013221-3212230112312322-2330320300130301)
- [cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password](resources--k8s_cluster--reference--group-001.md#canonical-3221232020303001-3003231303331233-1132210113332003-1211211323113310-3203312110221013-1332122020311003-0202033313310012-2313111303012003)
- cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info

<a id="canonical-1133210222130122-3120330223002320-0230230202233310-1230023023233011-2301121021131122-3131003001210102-2012021031021033-1122122031212002"></a>

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

<a id="canonical-2302003101120022-1033130020031201-0021210000321022-3211330210002320-1113321022021232-3310111130133303-3003001331221133-2202112121233203"></a>

### Direct properties for `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info`

<a id="canonical-1032002003100032-3030232203011033-1321321232120200-0303112220012320-2312032220110200-3011110213101230-2222131130322220-0102320011301120"></a>

#### `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0110333030330122-1021322010022221-2103100000232022-3232032332020002-1201133322030202-0202123111033020-0103333113122020-1202312213020301"></a>

<a id="canonical-3103111321133230-1231103333013321-3210320021132111-3310322123111112-0103120111010011-2331001331332331-2030101110112031-2113312132121333"></a>

#### `cluster_wide_app_list.cluster_wide_apps.argo_cd.local_domain.password.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3111130013310203-2021310113320211-1112100020113323-3123101301130332-1301211221103223-0123110011020110-0002210120232023-3010031222033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps.dashboard` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020)
- cluster_wide_app_list.cluster_wide_apps.dashboard

<a id="canonical-2232032312201122-3213220312121003-1023011133202202-2230003110200201-3303231003223133-3023220113222003-3002330213013332-3220333100320131"></a>

Type: `["object", {}]`. Optional.

Description Parameters for K8s dashboard.

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
dashboard = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031011133333033-1333111103030230-3032320110232222-3122022012122220-1202203013332031-2033031112321320-1311302221323121-2220010113000321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps.metrics_server` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020)
- cluster_wide_app_list.cluster_wide_apps.metrics_server

<a id="canonical-1000230123303313-0313331023011132-3003312220203332-1120231202221030-0103113133323001-0322333222102321-1123212003122121-2021100021303013"></a>

Type: `["object", {}]`. Optional.

Description Parameters for Kubernetes Metrics Server application.

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
metrics_server = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102233013113032-1120130222210211-2313102122223131-0212311203003303-2121032202130330-3000012030231030-2300301033202210-3023131302301013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cluster_wide_app_list.cluster_wide_apps.prometheus` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [cluster_wide_app_list](resources--k8s_cluster--reference--group-001.md#canonical-1323201222033222-1310002230121303-0210321122002030-0032003303033022-0012010230001210-3332312220000031-2123030011032302-3301012023021030)
- [cluster_wide_app_list.cluster_wide_apps](resources--k8s_cluster--reference--group-001.md#canonical-0200001310311212-1121331310012002-3033011130202120-1023002011031303-0200330313203121-0103121103310212-0321232101200102-0333031333312020)
- cluster_wide_app_list.cluster_wide_apps.prometheus

<a id="canonical-1230310330313303-3230203220201201-0120312220113120-1313223300331131-3221320130000001-0002032210110122-0203320211211132-0223213002001110"></a>

Type: `["object", {}]`. Optional.

Description Parameters for Prometheus server access.

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
prometheus = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110033101330033-3012032322223221-3012123023320130-1102100010210112-1112321122232211-0033011130032123-1102210111203133-1110010101001011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `global_access_enable` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- global_access_enable

<a id="canonical-0213110233322333-0223333313210033-2200331030213133-3221131213003323-1322122010031300-0120010011230331-1122030002300300-3132031023200210"></a>

Type: `["object", {}]`. Optional.

\[OneOf: global\_access\_enable, no\_global\_access; Default: no\_global\_access\] Configuration
parameter for global access enable.

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

- [global_access_enable](resources--k8s_cluster--reference--group-001.md#canonical-0213110233322333-0223333313210033-2200331030213133-3221131213003323-1322122010031300-0120010011230331-1122030002300300-3132031023200210)
- [no_global_access](resources--k8s_cluster--reference--group-001.md#canonical-3313313210202332-2111101111001300-0111010313003003-3122222132131112-3120213213332121-3113331200233203-0112110330033330-3110332002001012)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
global_access_enable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133012030210030-0220221330213223-2132132333211220-3120030220103203-1313003032301101-1300133311312022-1321021211021222-3230112201130332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `insecure_registry_list` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- insecure_registry_list

<a id="canonical-2100022003302111-1011113003200323-0001000100202102-3201032300101100-0133302230230000-0301133020103022-0121331223222130-2331233313210210"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: insecure\_registry\_list, no\_insecure\_registries; Default: no\_insecure\_registries\]
Docker Insecure Registry List. List of Docker insecure registries.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("insecure_registries")}
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

- [insecure_registry_list](resources--k8s_cluster--reference--group-001.md#canonical-2100022003302111-1011113003200323-0001000100202102-3201032300101100-0133302230230000-0301133020103022-0121331223222130-2331233313210210)
- [no_insecure_registries](resources--k8s_cluster--reference--group-001.md#canonical-3321121321223302-1210011303220123-3032020300021012-1220022120321133-2222201132113201-0310002312231202-1303320031013001-1211100113132101)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
insecure_registry_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200113003220312-2321200001021103-1102123103020120-3033113210313230-3130230100113232-2200103222133221-1121111100123201-2123212231213002"></a>

### Direct properties for `insecure_registry_list`

<a id="canonical-0012112211200110-1210123133001133-0113030332200132-0312301210200120-1100233121233123-3130320022223321-2232232212101130-0333233233301121"></a>

#### `insecure_registry_list.insecure_registries` property

Type: `["list", "string"]`. Optional.

List of Docker insecure registries in format 'example.com:5000'.

Additional upstream details:

List of Docker insecure registries in format "example.com:5000"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2330113202310232-1313112231102321-0100021023331021-3030000310021310-0120111130300032-1301121303202001-0330002111131003-3333210313223011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_access_config` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- local_access_config

<a id="canonical-2103201032303300-1320122020201231-0303123212022321-0132331230310013-2223322213330222-2132131001211030-0103132110110301-0311100111033223"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: local\_access\_config, no\_local\_access; Default: no\_local\_access\] Parameters required
to enable local access.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("local_domain"),
  validators.ConflictingObjectAttributes("default_port",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

OneOf alternatives in this subsection:

- [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-2103201032303300-1320122020201231-0303123212022321-0132331230310013-2223322213330222-2132131001211030-0103132110110301-0311100111033223)
- [no_local_access](resources--k8s_cluster--reference--group-001.md#canonical-3110122132203203-0211212131133312-3331101212003103-0320210130101012-2211012100230112-2130110102310220-2303213102030221-2101020110023220)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
local_access_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223111302303303-1331003230232200-0310103210002312-3031201313032303-3002311100020232-0230301111321101-2312231310002110-3211333233023023"></a>

### Direct properties for `local_access_config`

- [default_port](resources--k8s_cluster--reference--group-001.md#canonical-2203333201103021-0122002320130220-3131113321302033-3312332231312001-0111130130301101-3301101203103331-1120231312033310-2301323113210212): complete subsection reference.

<a id="canonical-3300111003312031-0131021221003110-2001203120302010-3211130331310032-1131210032102111-1320001020320322-3220232020322100-3102312112010222"></a>

<a id="canonical-2003013303111301-0231231220103002-3030031222032123-0023323031201330-3122121103033303-1011320312013312-3331332031110322-0212323133132110"></a>

#### `local_access_config.local_domain` property

Type: `"string"`. Optional.

Local K8s API server will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3010202023123233-1201310032301120-0030303113022002-2002330211232020-2221010320311103-1222232331333320-1111020021230031-2310013013011232"></a>

<a id="canonical-2200021031311123-3302213300023333-3213013330001301-1301202132332202-1300211130322032-1122130211230221-2122132232022022-2221031010321000"></a>

#### `local_access_config.port` property

Type: `"number"`. Optional.

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  }
}
```

<a id="canonical-2203333201103021-0122002320130220-3131113321302033-3312332231312001-0111130130301101-3301101203103331-1120231312033310-2301323113210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `local_access_config.default_port` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [local_access_config](resources--k8s_cluster--reference--group-001.md#canonical-2330113202310232-1313112231102321-0100021023331021-3030000310021310-0120111130300032-1301121303202001-0330002111131003-3333210313223011)
- local_access_config.default_port

<a id="canonical-1011003110233320-2210113003110302-1333111133231213-0203120302002012-1011022213120223-1301211232120322-1332210223321120-2333213101123231"></a>

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
default_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210131033311120-1312231122121133-3230223112010200-2121132220123123-1121312322203212-2333231201000112-3202301221212121-1221201220201322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_cluster_wide_apps` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- no_cluster_wide_apps

<a id="canonical-2301311322321010-0011123230311100-2331130103233323-2121301212012303-3211131023133301-0011123330122231-1220100310213211-3302233100001212"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_cluster_wide_apps = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102111002233003-1113331000033312-3121001211313133-2221120110302331-1321010313323021-0312101222012023-3103201132033230-1022003113010030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_global_access` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- no_global_access

<a id="canonical-3313313210202332-2111101111001300-0111010313003003-3122222132131112-3120213213332121-3113331200233203-0112110330033330-3110332002001012"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no global access. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_global_access = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032123322021013-3102121121302220-1301312300210220-0031131220103302-1310012310132313-3330231103033230-2313301123001103-3333332000300300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_insecure_registries` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- no_insecure_registries

<a id="canonical-3321121321223302-1210011303220123-3032020300021012-1220022120321133-2222201132113201-0310002312231202-1303320031013001-1211100113132101"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_insecure_registries = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000131320001322-3231301312222232-0321232331030302-1323131212032120-1021233030121121-0012033021120311-3332231013132102-3210020112130130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_local_access` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- no_local_access

<a id="canonical-3110122132203203-0211212131133312-3331101212003103-0320210130101012-2211012100230112-2130110102310220-2303213102030221-2101020110023220"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no local access. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_local_access = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222111132220221-2003132131331103-2132032010321023-0230211030303331-1333110222213201-2121103131130022-3021231323211331-2010330002322221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- timeouts

<a id="canonical-3210331111102101-2212302231321020-0023102223223201-0133032032332121-1331301302010010-1111232011232110-1132130313231020-2132230232101330"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102130313331011-2201013030001230-0220211211202032-2131313111230210-0300000130223120-3203332021102223-0202302230332122-0303303122213133"></a>

### Direct properties for `timeouts`

<a id="canonical-0120300022120030-0023201020020203-1120302213321233-1321220200102032-3123333302110201-1312031023131310-2102033110320202-0210212032231100"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1322331131212113-0132003002013200-3013011111010101-1112132213122113-2022120220200220-2122132232300133-1030211233003013-0332231002210213"></a>

<a id="canonical-2112103021203000-3012313113333002-1301102023021022-0133203003311001-0313233120103021-3312330002303122-3033313213331311-3111110222102120"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1312130332312022-1232131323031203-0212230123223000-2320230333321320-2230301211333103-0000213230103110-3230230121312123-0110132300003012"></a>

<a id="canonical-2032102212130200-3312131212011002-3002210031033102-2101231302102212-1131233301033101-1200012302300030-1111210311220321-2132103231231021"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2003122221200031-1030011311220003-3202220313201332-1312211312330103-0023311311330120-0111130122233130-3010122212002023-2033202310300231"></a>

<a id="canonical-2333211330330310-0322111030212303-1112312122232231-0001033311223023-1021123323333312-1033020212213122-2303221022123012-2301012101321203"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0211302322000100-0203322011230121-0312133132101320-1031210001131331-1031230221102200-3221230333033330-3133111123333321-0133230211020211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_custom_cluster_role_bindings` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- use_custom_cluster_role_bindings

<a id="canonical-1021121302311233-3033003300301332-2021111222212201-3033121221133222-3110101231113213-1133002131200303-2030303112330200-2000330133330122"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_cluster\_role\_bindings, use\_default\_cluster\_role\_bindings; Default:
use\_default\_cluster\_role\_bindings\] List of active cluster role binding list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_role_bindings")}
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

- [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-1021121302311233-3033003300301332-2021111222212201-3033121221133222-3110101231113213-1133002131200303-2030303112330200-2000330133330122)
- [use_default_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-2333331010123001-0300022310113220-2120001233210003-0213233233310313-1133030023130213-1311312002033201-1020133232011102-3111311012213101)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_cluster_role_bindings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112020132001330-0331300131232030-2211303310020221-3121021031312303-1213130113012000-2022323031013221-3303122111123323-1220112302033122"></a>

### Direct properties for `use_custom_cluster_role_bindings`

- [cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-1222330001100231-3313330110330313-2103130133020322-2313223033001121-1201120302332223-0103303210300201-2130220101032213-0303223032033213): complete subsection reference.

<a id="canonical-1222330001100231-3313330110330313-2103130133020322-2313223033001121-1201120302332223-0103303210300201-2130220101032213-0303223032033213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_custom_cluster_role_bindings.cluster_role_bindings` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [use_custom_cluster_role_bindings](resources--k8s_cluster--reference--group-001.md#canonical-0211302322000100-0203322011230121-0312133132101320-1031210001131331-1031230221102200-3221230333033330-3133111123333321-0133230211020211)
- use_custom_cluster_role_bindings.cluster_role_bindings

<a id="canonical-1311103030322321-2301132231011212-3122022103002011-1100302033021120-2323211213022101-1020112323022111-1112313312312111-3323002322133231"></a>

Type: `"object"`. list nested block, Optional.

List of active cluster role binding list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
cluster_role_bindings {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000230022000110-1200220203010113-3312010310332201-0331033200010033-0020233122022021-2020332130101310-2132111111030132-3002101300000322"></a>

### Direct properties for `use_custom_cluster_role_bindings.cluster_role_bindings`

<a id="canonical-0203102312330303-3333230223223313-0113233031102101-2333223001303223-1232022113012032-0330202012311120-0032222321231032-1000131032133222"></a>

#### `use_custom_cluster_role_bindings.cluster_role_bindings.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3121101103122021-3311132023113021-3123221333233232-1030022331332202-1310312030033012-1122211031100033-0303030010003011-2212332300313221"></a>

<a id="canonical-1011130002121312-2300111122201200-1022111321013233-2230230312132011-3233113321101233-1032221302221210-3313332032022030-0012320203031012"></a>

#### `use_custom_cluster_role_bindings.cluster_role_bindings.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0301203123102130-0231003020330033-1000013311211323-0100123031232333-0302120220112303-1202322131320230-0231300000123103-2012303333320333"></a>

<a id="canonical-2002202310010230-1022313331031121-1300001113202122-3222123203122023-1102221210302312-1200120330231121-1233020213032210-3000233020021120"></a>

#### `use_custom_cluster_role_bindings.cluster_role_bindings.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2033312333111003-3123320312002221-3221201213321321-0213023322032121-3232230123201131-0222231311213030-3330211220123210-0223222330320322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_custom_cluster_role_list` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- use_custom_cluster_role_list

<a id="canonical-0322303213323223-3200020033032230-3000200200031113-2021020211322120-3313230231121231-1300231021300120-3311030230302222-3302201212112032"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_cluster\_role\_list, use\_default\_cluster\_roles; Default:
use\_default\_cluster\_roles\] List of active cluster role list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_roles")}
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

- [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-0322303213323223-3200020033032230-3000200200031113-2021020211322120-3313230231121231-1300231021300120-3311030230302222-3302201212112032)
- [use_default_cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-3321330210122120-1033210002331112-0300110101203330-3230233323313033-2032012001212121-1203031310100031-2002120323003110-2332013121311321)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_cluster_role_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132131213301210-2000212332212303-3033211031100033-3101303013200210-0232213213321023-0302012210332201-0223203302331130-0033320302333302"></a>

### Direct properties for `use_custom_cluster_role_list`

- [cluster_roles](resources--k8s_cluster--reference--group-001.md#canonical-0330010311013002-3211232211212002-1130322333202333-3101012302321222-0032212233112303-3022031020230120-1001132103232322-0122300233002221): complete subsection reference.

<a id="canonical-0330010311013002-3211232211212002-1130322333202333-3101012302321222-0032212233112303-3022031020230120-1001132103232322-0122300233002221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_custom_cluster_role_list.cluster_roles` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [use_custom_cluster_role_list](resources--k8s_cluster--reference--group-001.md#canonical-2033312333111003-3123320312002221-3221201213321321-0213023322032121-3232230123201131-0222231311213030-3330211220123210-0223222330320322)
- use_custom_cluster_role_list.cluster_roles

<a id="canonical-0222330201030200-1130023021002111-0210311020013310-3013113233012111-1303003311202112-0113010332322022-2031132232130211-3322022311311230"></a>

Type: `"object"`. list nested block, Optional.

List of active cluster role list for a K8s cluster.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
cluster_roles {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131211033301221-3123100323011302-1332033112020003-2033320211202001-1010113213232113-0321133332133122-3221012230230210-1223000213012132"></a>

### Direct properties for `use_custom_cluster_role_list.cluster_roles`

<a id="canonical-1031302202112211-1211012023213132-0022233103321222-3023312011020101-1133322113200220-1201211101210100-3011310012021103-3233232103122301"></a>

#### `use_custom_cluster_role_list.cluster_roles.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0210223200023213-2111332212101103-1231132121322311-1122302213121100-3220223132012131-1123300031301112-3330320031022122-1023010033013123"></a>

<a id="canonical-2111322110202200-2003111202130212-1323322022002211-2321210310230132-0103322210100313-3210231230100000-0303022332002231-1023231103310011"></a>

#### `use_custom_cluster_role_list.cluster_roles.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3102233333010133-2232331001031333-0022032122203123-1333230113122212-1312021110031031-0130212200100303-3112213332330033-2203210313122113"></a>

<a id="canonical-2231021111030120-1000011203021000-2010121312132312-0300202303221333-1130002032303321-1302222032233323-1122200222233032-3323230201101312"></a>

#### `use_custom_cluster_role_list.cluster_roles.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2312030321230203-3200103233112030-1101003101310211-1303122131333011-3312132231111332-1031233123320221-1203002202101313-2123320021313222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_custom_pod_security_admission` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- use_custom_pod_security_admission

<a id="canonical-3300233020112301-2311123202333232-2031121320200311-2313313022022230-0212210211003011-3211022120222213-1320213322100212-3231131212202103"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_pod\_security\_admission, use\_default\_pod\_security\_admission; Default:
use\_default\_pod\_security\_admission\] Type establishes a direct reference from one object(the
referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

- [use_custom_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-3300233020112301-2311123202333232-2031121320200311-2313313022022230-0212210211003011-3211022120222213-1320213322100212-3231131212202103)
- [use_default_pod_security_admission](resources--k8s_cluster--reference--group-001.md#canonical-0100202212033302-2322133002120211-0231000210210323-0101012210103120-1320101031313213-2132222030100232-2103020200102333-1211300022221301)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_pod_security_admission {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033333011300113-2012301232223311-1332100001211102-2021121002133321-0121132323030203-2131313330213133-1222001233131022-1121112100101112"></a>

### Direct properties for `use_custom_pod_security_admission`

<a id="canonical-1313032303103201-3322302221232223-0333201122010130-0003123003321030-2201233220021231-2320322210213120-1211022322131323-2021021021012110"></a>

#### `use_custom_pod_security_admission.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2000101231132133-0010320312323122-2022111233313212-3210132223330330-2020111222221023-1021332033002300-2303130230201213-2002332133303303"></a>

<a id="canonical-0210201120131322-1133020303222321-3020201320323012-3111001331303000-3120031200222332-0100210113322302-0011332120131311-3303103131010311"></a>

#### `use_custom_pod_security_admission.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3331032112211123-0320012323301322-3100112112103010-1002111131322303-2101030101023120-1232221002023303-2302211113112311-0313121310223312"></a>

<a id="canonical-2101113120120221-1101212102120333-1011013322003222-1332212102331212-2023210130302220-3023101332002033-1233330331031120-3002330100221203"></a>

#### `use_custom_pod_security_admission.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3131020221100132-1003100033303012-3130330301313330-1033003233300201-1331103000203002-0033021330203230-3013312030313301-0011311101333322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_custom_psp_list` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- use_custom_psp_list

<a id="canonical-1131122010200101-0301221002021212-2031120220103122-1220101122031233-3030210323122033-1032313223020321-0131221300120122-3103332022210303"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: use\_custom\_psp\_list, use\_default\_psp; Default: use\_default\_psp\] List of active Pod
security policies for a K8s cluster.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("pod_security_policies")}
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

- [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-1131122010200101-0301221002021212-2031120220103122-1220101122031233-3030210323122033-1032313223020321-0131221300120122-3103332022210303)
- [use_default_psp](resources--k8s_cluster--reference--group-001.md#canonical-2330110111230123-1223231130133333-2011031001010130-3231111012320030-0023322012102020-1003002331130312-1130001312131101-2331023211112012)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
use_custom_psp_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320311310110131-3330300100223321-1121300123131001-0321120003013213-2010222130002002-3312330123333031-3320320223122302-3322121130103202"></a>

### Direct properties for `use_custom_psp_list`

- [pod_security_policies](resources--k8s_cluster--reference--group-001.md#canonical-0212111212212021-3011323230102122-1202310001220202-2301303321012313-3012020123012013-1311002120232123-2113012331010033-0321020003333003): complete subsection reference.

<a id="canonical-0212111212212021-3011323230102122-1202310001220202-2301303321012313-3012020123012013-1311002120232123-2113012331010033-0321020003333003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_custom_psp_list.pod_security_policies` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- [use_custom_psp_list](resources--k8s_cluster--reference--group-001.md#canonical-3131020221100132-1003100033303012-3130330301313330-1033003233300201-1331103000203002-0033021330203230-3013312030313301-0011311101333322)
- use_custom_psp_list.pod_security_policies

<a id="canonical-2322203130113112-1101133031321302-0301222330212123-2113003102221210-0211312302311300-1202203213200320-1311222100212132-2333232111330320"></a>

Type: `"object"`. list nested block, Optional.

List of active Pod security policies for a K8s cluster.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
pod_security_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032003012003331-1310032331301132-1221020110332222-0103200110330312-3331032031302221-1312213212210332-3121021302000013-2002321300010001"></a>

### Direct properties for `use_custom_psp_list.pod_security_policies`

<a id="canonical-1101032010301010-2213131112302211-2312313211231221-0302032301121312-3120222221120300-0110011221233330-1202220222131223-2312213323303312"></a>

#### `use_custom_psp_list.pod_security_policies.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3112211333030212-2332300023013002-0120132110230233-3011312200113133-3320111231233332-3020132030222303-1013223310311210-2100122030230313"></a>

<a id="canonical-2123033201000232-2200113331100320-3101011221111203-0102331120201200-3320330300320123-2022233031030322-3020031123031203-0313310121132102"></a>

#### `use_custom_psp_list.pod_security_policies.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2311132231222210-3201120210023112-3211022213033002-0112312111101301-1110000011100300-2002120223033320-1123132201132223-0031201312221120"></a>

<a id="canonical-2233300203123022-0130101003103333-1122022123011202-0213230030322221-3303322201022003-3231302032110223-0110312123010223-2213012122320321"></a>

#### `use_custom_psp_list.pod_security_policies.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3023013200022221-2131212202021103-2232231332300220-1010102031222320-2310231313121132-0120123013030132-1233103103321123-1213221201311211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_cluster_role_bindings` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- use_default_cluster_role_bindings

<a id="canonical-2333331010123001-0300022310113220-2120001233210003-0213233233310313-1133030023130213-1311312002033201-1020133232011102-3111311012213101"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
use_default_cluster_role_bindings = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213233222332312-2010100313000111-1130013031231223-0203332222133032-1321023231023111-0132310322222130-0213113132211321-0130123032022033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_cluster_roles` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- use_default_cluster_roles

<a id="canonical-3321330210122120-1033210002331112-0300110101203330-3230233323313033-2032012001212121-1203031310100031-2002120323003110-2332013121311321"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
use_default_cluster_roles = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302100303223331-0223211122120012-2020001102301202-0333121332233033-0310133202333012-0111223110010210-1212202201010123-3322101332123102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_pod_security_admission` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- use_default_pod_security_admission

<a id="canonical-0100202212033302-2322133002120211-0231000210210323-0101012210103120-1320101031313213-2132222030100232-2103020200102333-1211300022221301"></a>

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
use_default_pod_security_admission = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301302001120301-3120213203003133-3013101332200030-3303203223000211-1100230021000103-3230002012102133-0311030302011303-2322000302222001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_psp` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- use_default_psp

<a id="canonical-2330110111230123-1223231130133333-2011031001010130-3231111012320030-0023322012102020-1003002331130312-1130001312131101-2331023211112012"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for use default psp. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
use_default_psp = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210013331002303-1110100111123032-0010132331001203-0120302213302320-1110231110012230-3113330311303112-3232113213313233-1023013222120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vk8s_namespace_access_deny` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- vk8s_namespace_access_deny

<a id="canonical-1230112122021332-1321330221320110-3030121021021303-3032313030210320-3231101112320131-0302120333312322-3012310121323302-1322021020212110"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: vk8s\_namespace\_access\_deny, vk8s\_namespace\_access\_permit\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [vk8s_namespace_access_deny](resources--k8s_cluster--reference--group-001.md#canonical-1230112122021332-1321330221320110-3030121021021303-3032313030210320-3231101112320131-0302120333312322-3012310121323302-1322021020212110)
- [vk8s_namespace_access_permit](resources--k8s_cluster--reference--group-001.md#canonical-1210200220020311-0221232223021310-2011020012103310-1111201321220202-3010002102211322-0100312231011033-1322210023021220-0131212330010113)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
vk8s_namespace_access_deny = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011031013012102-2201320113333031-2201203010203212-0000332201221301-1301311203111302-2022113130022030-1222210102113232-3302123213330233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vk8s_namespace_access_permit` properties

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md#canonical-0233201000200102-2023223010011110-3102210330122033-0203123200303231-2213002033003103-0222302331113011-3300321101303322-0301121023303212)
- [Property reference](resources--k8s_cluster--reference--group-001.md#canonical-2100330132213030-1101211002122201-2110222313321022-3031232100331321-0130001220020011-3313222233113001-0323333221302213-3332212002101312)
- vk8s_namespace_access_permit

<a id="canonical-1210200220020311-0221232223021310-2011020012103310-1111201321220202-3010002102211322-0100312231011033-1322210023021220-0131212330010113"></a>

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
vk8s_namespace_access_permit = {}
```

This is an empty object or choice marker. It has no direct properties.
