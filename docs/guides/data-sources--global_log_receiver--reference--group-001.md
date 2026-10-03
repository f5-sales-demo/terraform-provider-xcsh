---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301302212001110-2322313121323103-0020222221011330-1230300312311031-0113310122322211-0032333321011302-3212203000221230-2122002322200111"></a>

## Property reference — Property reference / 220322330022 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- Property reference

<a id="canonical-2320211021020003-3330113210223201-1331311231302223-1221131023332231-2031103013102221-2213303021101221-0102211030011320-1002200310220312"></a>

## Direct properties — Property reference / 220322330022 / 3

<a id="canonical-0020011300323223-1311001020021023-2030212233021202-3223323332113331-2010232113000000-1222310201110301-2000020201031012-3013320110021020"></a>

<a id="canonical-1101111013133001-3200212123320132-2133200003313320-2131200333011110-3331332003130112-2300011103311211-3102303030331001-3032330110103022"></a>

## annotations property — Property reference / 220322330022 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

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

- [audit_logs](data-sources--global_log_receiver--reference--group-001.md#canonical-3133212020032302-0323212203031203-3200200231033010-2303311030030012-3001102312131320-3233231303203033-1101020321312133-2201321333233323): complete subsection reference.

- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032): complete subsection reference.

- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-0132210313021033-2221133303001132-2002201122230200-2203313221112330-3122130230230331-3110223323220223-1131013312112132-1123200222220321): complete subsection reference.

- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-2121301220330030-1221312033213003-0320321302002303-2133212312111123-3330122101022123-0133210113001032-2201030003210013-3211030023023312): complete subsection reference.

- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102): complete subsection reference.

<a id="canonical-1210131003031200-1302123200313011-3303320111231020-1030311102301232-2201122030001101-1230320233100002-0333010312223231-0233010330302113"></a>

<a id="canonical-3011311120310131-3230202121030331-0022232330212001-2001013132113311-1321322102233131-0023330200332031-2213133003322201-1111201122201323"></a>

## description property — Property reference / 220322330022 / 5

Type: `"string"`. Computed.

Description of the GlobalLogReceiver.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [dns_logs](data-sources--global_log_receiver--reference--group-002.md#canonical-2313012101200221-3103103122211233-0010031133112232-3010013110321112-0132120212211113-2313113031330310-0333130212030033-3300322103021000): complete subsection reference.

- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302): complete subsection reference.

- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220): complete subsection reference.

<a id="canonical-2200111230003002-2201231231122133-1301001313221032-0012331210330233-1303000023223013-0223130312133233-0210312022013301-1120323033330321"></a>

<a id="canonical-3210132102311220-2331331223320102-0122301000011201-0121203101212000-0103122203333030-3002331220330320-0311230300230203-2200120212312221"></a>

## ID property — Property reference / 220322330022 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331): complete subsection reference.

<a id="canonical-2110212111112102-1220100112002313-3122230022332000-0313223100220331-1100133322321120-2321120232332320-2231113132032002-3123011133021003"></a>

<a id="canonical-0212123220013113-2001122001223032-1300110202322130-2320030131132300-3112032022213033-1100113012220030-0030031111223332-1223123231321031"></a>

## labels property — Property reference / 220322330022 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

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

<a id="canonical-0000131312002213-1233112220201021-2002310231232330-2230202330211020-3222221223213313-3101121233223132-0110013201123103-3201000102330023"></a>

<a id="canonical-1331102033201111-0010110100131231-3002132121032223-2011120220331303-3130311303232213-1310101011130120-1322310313101130-0001120330200020"></a>

## name property — Property reference / 220322330022 / 8

Type: `"string"`. Required.

Name of the GlobalLogReceiver.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1011201023030032-2202132102310112-1121133112301202-1203123302021102-1330333111131013-0233010133013200-0020201331003133-3200230012123121"></a>

<a id="canonical-2020132311222330-2101213202132110-2103021310212310-1102010020023310-2023133332021121-3320110330313313-1221113333300200-2213310231211211"></a>

## namespace property — Property reference / 220322330022 / 9

Type: `"string"`. Required.

Namespace where the GlobalLogReceiver exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-1222203022011122-1213111132203013-1201312120332200-3233301331202013-3012133313300230-3320033231321122-3201112101101230-3232031212020130): complete subsection reference.

- [ns_all](data-sources--global_log_receiver--reference--group-003.md#canonical-3222301230023030-0001112112122320-3100220123111200-3201120222111311-1330032121023233-1201222312320311-1111331322201210-0133312301312020): complete subsection reference.

- [ns_current](data-sources--global_log_receiver--reference--group-003.md#canonical-3213312020222012-2002220201233002-0031302031132022-2311120330010033-0031032202323202-0022102301032121-0321103001233122-0001120123201310): complete subsection reference.

- [ns_list](data-sources--global_log_receiver--reference--group-003.md#canonical-3303120033132301-0202221321231331-0130123312332131-0202000200222232-3021322301011113-1000013103130021-2133210101223030-0302330311022122): complete subsection reference.

- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203): complete subsection reference.

- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123): complete subsection reference.

- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202): complete subsection reference.

- [security_events](data-sources--global_log_receiver--reference--group-004.md#canonical-1002031033031310-1233001100003213-2003102123102332-1202233033221123-1021111002311202-0121010220310110-1211132100130223-2001223323310023): complete subsection reference.

- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300): complete subsection reference.

- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001): complete subsection reference.

<a id="canonical-1112231020332222-1300222000313220-1002003012313313-0332311221322000-1321222200000113-2223120032021311-1213133312310223-3112331023200033"></a>

## All schema paths — Property reference / 220322330022 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--global_log_receiver--reference--group-001.md#canonical-0020011300323223-1311001020021023-2030212233021202-3223323332113331-2010232113000000-1222310201110301-2000020201031012-3013320110021020) |
| `audit_logs` | [audit_logs](data-sources--global_log_receiver--reference--group-001.md#canonical-2133221203001223-3123231200122131-3110232122221212-2033103202301212-0132123303320010-3111333302033122-1110213131202301-2111222131331002) |
| `aws_cloud_watch_receiver` | [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-3331110112131012-1000100132111012-1203111230222102-1230112233112113-0001200333200231-0312301323230303-1011311201201322-2301312200332113) |
| `aws_cloud_watch_receiver.aws_cred` | [aws_cloud_watch_receiver.aws_cred](data-sources--global_log_receiver--reference--group-001.md#canonical-3121101011123321-1230031310012031-3011202211032002-0233302301311202-1232122303112311-2232210033101011-0301200003030332-1300023211110232) |
| `aws_cloud_watch_receiver.aws_cred.name` | [aws_cloud_watch_receiver.aws_cred.name](data-sources--global_log_receiver--reference--group-001.md#canonical-0202130221020102-1002120220201021-2101131131032010-3300122130203203-2333101121121223-1332323121320200-0020000012333230-3233013202320221) |
| `aws_cloud_watch_receiver.aws_cred.namespace` | [aws_cloud_watch_receiver.aws_cred.namespace](data-sources--global_log_receiver--reference--group-001.md#canonical-1103211111032210-1321331101303301-1200323312010012-0303013330012030-1002222032321301-1101001013112323-0132232132321121-3110203220300033) |
| `aws_cloud_watch_receiver.aws_cred.tenant` | [aws_cloud_watch_receiver.aws_cred.tenant](data-sources--global_log_receiver--reference--group-001.md#canonical-3003003123111310-3301010313313211-1012230001230313-2031032002322300-3132030211313020-2121001111202112-3132333113210113-0322331330320201) |
| `aws_cloud_watch_receiver.aws_region` | [aws_cloud_watch_receiver.aws_region](data-sources--global_log_receiver--reference--group-001.md#canonical-3311102200012103-3032303221112130-0113102011002001-2133021101120132-1110010032220322-3312010332100033-1320120110331131-1233212302121121) |
| `aws_cloud_watch_receiver.batch` | [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-1233331313232013-0210102231321101-2233132230112030-1102102102322100-0210030231312322-2131002102132002-2121002032120031-0113012201312130) |
| `aws_cloud_watch_receiver.batch.max_bytes` | [aws_cloud_watch_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-001.md#canonical-3111222320110330-0211233113013301-0130332330113333-0332103102301221-3113032313132223-0003223213322131-3333031300120331-0133313030310133) |
| `aws_cloud_watch_receiver.batch.max_bytes_disabled` | [aws_cloud_watch_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-1310033223231011-2332311222312102-1321322133332200-3333101123123000-1210330113013203-1233031223120101-1010210133222321-1120130101322321) |
| `aws_cloud_watch_receiver.batch.max_events` | [aws_cloud_watch_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-001.md#canonical-2130002032131222-2211200213033111-2232010310021320-1310010021110130-1232311111232010-2311131202213023-3002132101110331-3000211222203220) |
| `aws_cloud_watch_receiver.batch.max_events_disabled` | [aws_cloud_watch_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-2302110232231000-2323222320201120-1202001230202313-3033132222223122-1012331302331033-0230002113201223-0333132033313100-2113020031213232) |
| `aws_cloud_watch_receiver.batch.timeout_seconds` | [aws_cloud_watch_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-001.md#canonical-0020211302321221-2113321213220001-1030001002023023-1223112032202120-1033230003020322-2000022013013333-3121121103110132-2113023121100310) |
| `aws_cloud_watch_receiver.batch.timeout_seconds_default` | [aws_cloud_watch_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-001.md#canonical-2312120211222333-2032311303102131-3000122003213310-2023032013321331-3131121321113101-0131112102220231-2202003011330311-2031132130122100) |
| `aws_cloud_watch_receiver.compression` | [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-2110121120030013-2001320103312120-2103313323000223-0321121211112313-0101200213310000-2031202201132121-0223212033011110-0332311102312013) |
| `aws_cloud_watch_receiver.compression.compression_default` | [aws_cloud_watch_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-001.md#canonical-1030032322103222-1133311133112000-2303221023222200-2023221031233320-2331300230013000-3113200330130132-2212202131211123-3103123212133302) |
| `aws_cloud_watch_receiver.compression.compression_gzip` | [aws_cloud_watch_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-001.md#canonical-1223223221203132-0331322110213000-3202231011223233-3001001102323313-1213111220211133-3101313220122203-3233101230133023-3212122011020030) |
| `aws_cloud_watch_receiver.compression.compression_none` | [aws_cloud_watch_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-001.md#canonical-0012320110221120-2013320300301201-0011031230103110-2030221131121102-1332023310233231-3131321313331031-1202102231233210-1203032313302332) |
| `aws_cloud_watch_receiver.group_name` | [aws_cloud_watch_receiver.group_name](data-sources--global_log_receiver--reference--group-001.md#canonical-0013332233001311-3323103011023333-2313200223201223-3013332031232222-3330333211033331-0220120010222303-1303222023012103-0100221120021302) |
| `aws_cloud_watch_receiver.stream_name` | [aws_cloud_watch_receiver.stream_name](data-sources--global_log_receiver--reference--group-001.md#canonical-1022333110331123-1031010001031013-0332110031122222-1002112102020010-1330311002201211-1301311213222120-1032013031023220-0303332030021303) |
| `azure_event_hubs_receiver` | [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1133203103010030-0213101233221030-0113301022130000-0112220222113332-2113012101223231-1210310130311121-1322212023012203-3120213302333232) |
| `azure_event_hubs_receiver.connection_string` | [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-0333100212131103-1133011001333101-2212011112023001-3200022001301210-1200012031211302-2001013303132330-1333301001202001-3101030311232032) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-1233121333012322-0212112222222323-1330121130103201-0232112020111102-0220213113221220-0100102010110321-1100302120020220-3101131233210330) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-001.md#canonical-2131033323211322-0200320002201333-3020130303303032-1332212123120111-2310032020200013-0112121130023211-2223210120322301-3020101111030133) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.location` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-001.md#canonical-2221123102131101-3003030122112302-1322023112122132-2020330332121133-1332112033222232-2231200322310331-1100231322230011-0131103023113132) |
| `azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_event_hubs_receiver.connection_string.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-001.md#canonical-0202201213021030-1313123303303022-3130220112012123-1302032232101313-3223231113110010-2200311112031022-1312022002221300-3233022311102233) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info` | [azure_event_hubs_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-1130233320221032-1321331300233201-1322133322211221-1333221021321010-0210232333033320-1313200022332122-3222320230232103-0012030210003332) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref` | [azure_event_hubs_receiver.connection_string.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-001.md#canonical-3331033000310013-2231323032112021-2013021210110332-3002131323323202-0312311013032012-1220303122100231-1110101012333022-0102003132333110) |
| `azure_event_hubs_receiver.connection_string.clear_secret_info.url` | [azure_event_hubs_receiver.connection_string.clear_secret_info.url](data-sources--global_log_receiver--reference--group-001.md#canonical-2103133030323122-3301103203000210-2101212302202333-0233033123231121-2101122311233323-0033010231203302-2302301133223322-0322203202201213) |
| `azure_event_hubs_receiver.instance` | [azure_event_hubs_receiver.instance](data-sources--global_log_receiver--reference--group-001.md#canonical-3233003101121100-2112231100120032-1012002220321100-0120103120021231-2301023112312020-2333231213300302-0211111022020323-2033221033211200) |
| `azure_event_hubs_receiver.namespace` | [azure_event_hubs_receiver.namespace](data-sources--global_log_receiver--reference--group-001.md#canonical-3212012100330311-0001100130100130-3323321211212101-0311033223300313-0033220010120001-0223010303230200-3012221313212321-1222232211330022) |
| `azure_receiver` | [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-0303310011011220-0200131310311211-2300302321112220-1311210110210130-2100302022132032-3210212112313131-3213323010012223-2203111201011331) |
| `azure_receiver.batch` | [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-2323221120203333-2211030001033223-0010121333110011-3201032021133322-3130332033033302-2322031012132313-1202021230002111-2223313330110213) |
| `azure_receiver.batch.max_bytes` | [azure_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-001.md#canonical-2322030301221021-2103122000031332-3322331030122111-2023200231111133-2330213331211012-0233023012023012-2221322302000233-0120120123001312) |
| `azure_receiver.batch.max_bytes_disabled` | [azure_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-3230320223211202-3112230121000212-1323312120200000-0222120003010130-2220103011120103-3213111103121303-2201223002001201-0031302133130121) |
| `azure_receiver.batch.max_events` | [azure_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-001.md#canonical-2220011103021002-0300030232032122-1112001113300003-3131102331313031-3303111312122130-2020223222013032-2030200122010033-0300031220132212) |
| `azure_receiver.batch.max_events_disabled` | [azure_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-0202312312231111-0122201100123102-0020001203123303-1131212022030331-2211033221112011-2030113112302001-0122330330010330-1111013201113112) |
| `azure_receiver.batch.timeout_seconds` | [azure_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-001.md#canonical-3110031303113112-1220203021023112-1322303032012133-2220100302233331-2122203221213033-1132330210231232-1223002221232220-2331101112123130) |
| `azure_receiver.batch.timeout_seconds_default` | [azure_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-1123030103200122-3233033132120003-0133122301003033-3313102012013210-1312310201103000-3301201322130222-3133122132033221-1031122131302023) |
| `azure_receiver.compression` | [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-3102130320331010-2111133122010001-2231313221032223-0203312101302211-2023300113023323-1030201321200220-0223101212020133-0000311213112113) |
| `azure_receiver.compression.compression_default` | [azure_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-0020120220203000-2012300302200002-2210202331010101-1121211311023232-1202300110121111-2020003123030312-1211103003312130-2300033302100311) |
| `azure_receiver.compression.compression_gzip` | [azure_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-2033001210123332-0122312030030311-1020321312231131-3001211003221322-2002120021300203-1221013211312000-1312013203021313-1332000311022020) |
| `azure_receiver.compression.compression_none` | [azure_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-3001303301131322-3020003013003220-1212300113012002-2333031022001232-3132222112213311-3001232121011132-0310231330032233-2113220032200112) |
| `azure_receiver.connection_string` | [azure_receiver.connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-2122202223022312-1310200121011323-1202122122122000-1130202102333300-2023031010111030-1211110312000332-2123013102202112-0221120200322313) |
| `azure_receiver.connection_string.blindfold_secret_info` | [azure_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-2123120222321013-2011001200110033-1301331333301323-3021310332101000-2031132232130000-3200011202120021-3213013233020232-0132020222201122) |
| `azure_receiver.connection_string.blindfold_secret_info.decryption_provider` | [azure_receiver.connection_string.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-2213101032233021-1111001000103131-1201313032221123-3312121313022321-0030321113310013-1232010111202133-2210031131303310-0021132023010001) |
| `azure_receiver.connection_string.blindfold_secret_info.location` | [azure_receiver.connection_string.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-2221222113333111-3313212220101102-0322000010212203-0102120121310012-3300123003313213-2331003302202312-2312322033320232-0030323103111220) |
| `azure_receiver.connection_string.blindfold_secret_info.store_provider` | [azure_receiver.connection_string.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-1110313312022231-2303120132312332-3230123332303113-2300222123120222-1211101013112123-1231123130003220-1101122012021213-0221111121012131) |
| `azure_receiver.connection_string.clear_secret_info` | [azure_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-2102123013221100-2231203100222021-1230321213131323-1113031031321003-0133230101203011-3233130232310032-3021222113100001-3120211121320132) |
| `azure_receiver.connection_string.clear_secret_info.provider_ref` | [azure_receiver.connection_string.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-002.md#canonical-3322302310102202-2132032102030021-1300220302331002-3001211131302300-0302031100331131-3023300323001002-2302232110233300-3033220321313111) |
| `azure_receiver.connection_string.clear_secret_info.url` | [azure_receiver.connection_string.clear_secret_info.url](data-sources--global_log_receiver--reference--group-002.md#canonical-2031223102332203-2213021112231031-0111020222022213-2320032310312123-3023121102122130-1323322011331022-3023000120313100-3300110000120230) |
| `azure_receiver.container_name` | [azure_receiver.container_name](data-sources--global_log_receiver--reference--group-001.md#canonical-1122122002230032-1001101212113032-0210102101211011-2133220013310021-2211211020032000-2012313301021132-1313321133323322-2002023220030010) |
| `azure_receiver.filename_options` | [azure_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-3103320202120000-3231111110312220-3033222232220233-3000202311220011-0222300222130203-2012132003113000-0213122020212211-1020222220000000) |
| `azure_receiver.filename_options.custom_folder` | [azure_receiver.filename_options.custom_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-0211210103230121-1123121320320130-0310120110133112-2121203301012310-3330132321033123-3110313132213000-1321312102320023-3323111201003331) |
| `azure_receiver.filename_options.log_type_folder` | [azure_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-2023020012222102-3100222112221022-2103033012133312-3030312310322223-3011032120103032-1210322020230301-1210303001220333-2210321030300203) |
| `azure_receiver.filename_options.no_folder` | [azure_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-2211300003231023-0221022012320100-2213112020100013-1112332101300003-1123222001031213-0130233203231323-3022322010332222-3130330030232000) |
| `datadog_receiver` | [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1030122021022232-1220312302320123-0120123303110130-0320113101000121-3332333032132102-1213013023033133-1200133200111110-3230023131300203) |
| `datadog_receiver.batch` | [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-2032303320333220-1010111220121000-3120101202200030-2303220030321331-2330002010311101-0010222032033221-3300013021300232-2311301300230111) |
| `datadog_receiver.batch.max_bytes` | [datadog_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-002.md#canonical-2201201300013211-1130123201032202-3033013131133212-2231013233030130-0012101030300122-0331201133230020-0301100331001202-0131010101323132) |
| `datadog_receiver.batch.max_bytes_disabled` | [datadog_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-0300120313120120-2202221113213322-3302232200330131-3210103022220310-2313110200331111-2313101032012301-1132221333303301-3121103020212330) |
| `datadog_receiver.batch.max_events` | [datadog_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-002.md#canonical-3031303110033232-3020311223201120-2221322331021332-1110100023000312-1301111232002132-0321001310120011-2120021321131111-3232030121112303) |
| `datadog_receiver.batch.max_events_disabled` | [datadog_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-1203231323222302-2303123302223011-3323210123033010-2331211032103030-3103133001011320-2101110233023000-1121322003211130-3033300333330131) |
| `datadog_receiver.batch.timeout_seconds` | [datadog_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-002.md#canonical-1121313132330311-3331030031303121-3002230102101032-3231333101101330-1310201232211033-3132113131323311-3133100132103112-3111011223323330) |
| `datadog_receiver.batch.timeout_seconds_default` | [datadog_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-2311113312202122-0030001323123122-2122022203000220-2131021032111203-0303230330203033-0232112023200132-2131323222033123-0312221000120233) |
| `datadog_receiver.compression` | [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-2111012231103312-1002111120000130-0123233201020312-3100212332012222-2000212013103302-3131111011113210-2333000321013312-1201003021302321) |
| `datadog_receiver.compression.compression_default` | [datadog_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-3210033122031313-2322331022312232-3111023321233320-1321212101202310-3323123103132333-1112313023101011-0132301221103301-1113011311233323) |
| `datadog_receiver.compression.compression_gzip` | [datadog_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-1003313101203011-0002121021222002-1131221011123121-0030333001332310-1311121311301031-3320120313022022-2002232330012113-2000022131221022) |
| `datadog_receiver.compression.compression_none` | [datadog_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-2032213033110020-3311133303310011-0233122210221130-0332302102122132-0311310132021012-3231310310020002-3133232023330331-0121231311011111) |
| `datadog_receiver.datadog_api_key` | [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-0323101312103313-0230000103212323-0210123202113212-2123013010210113-0231132122010111-3130122121223322-2022322330031131-0203321101313112) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info` | [datadog_receiver.datadog_api_key.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-3332200330132123-0323112202222013-1313202200232322-2113122320222202-0123020313013030-0010213322023021-3210032310113030-1011123330111230) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-0012132103310200-0311202103123212-0303010331133003-2303310101333101-3113210133022030-2100033203131232-0331003300211031-0332132112200002) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.location` | [datadog_receiver.datadog_api_key.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-0121032003133111-0222121113033312-3302300001221211-1032131110320231-0112031133030203-0213111331100111-2223102131000312-1213202100232300) |
| `datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider` | [datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-0331010232012313-3113200031122120-3301311322230013-3032023010330220-0321310311323130-0300012203203101-1100133300231221-2233112031211102) |
| `datadog_receiver.datadog_api_key.clear_secret_info` | [datadog_receiver.datadog_api_key.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-0133211001131000-0032320011101230-2020220311222112-2223201232332301-0200230312020303-3011003000322001-1311021320212112-3333102000303303) |
| `datadog_receiver.datadog_api_key.clear_secret_info.provider_ref` | [datadog_receiver.datadog_api_key.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-002.md#canonical-0021202032020031-0230130000322130-1023230121203001-0323020230101100-2030020130311000-3222111131121313-2213030121031230-2131221311031020) |
| `datadog_receiver.datadog_api_key.clear_secret_info.url` | [datadog_receiver.datadog_api_key.clear_secret_info.url](data-sources--global_log_receiver--reference--group-002.md#canonical-0120221103100211-1221122300302310-1231333321110210-3001202200003003-1312322132010322-3300223311332120-1013221110301102-0210221222002213) |
| `datadog_receiver.endpoint` | [datadog_receiver.endpoint](data-sources--global_log_receiver--reference--group-002.md#canonical-0302321030210011-1122020310101231-2023131212211203-1023021002302102-3113012103312021-3203100223211332-0122330032213100-0303301203121013) |
| `datadog_receiver.no_tls` | [datadog_receiver.no_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-2023313133203132-2232033120000100-2200002311023200-1232130322303112-0001332320232113-2030320120121102-1203123131022000-3120301030320321) |
| `datadog_receiver.site` | [datadog_receiver.site](data-sources--global_log_receiver--reference--group-002.md#canonical-1133000031211321-2030321133010031-3322103330312201-0030223220233312-0121112011322112-0313111033200130-3310320331323122-1122220233121032) |
| `datadog_receiver.use_tls` | [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-3103323002222123-0322201230323303-2323300302110331-0132103001133223-3000020103120200-1013233302203210-3230211020130010-2210013321113013) |
| `datadog_receiver.use_tls.disable_verify_certificate` | [datadog_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-3211012121013331-0301122231131101-2110120121103023-1011031101221000-1001021010120120-2103333003200111-1320102111222122-3301022311320223) |
| `datadog_receiver.use_tls.disable_verify_hostname` | [datadog_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-2331332232023132-2121000312003003-1331123103010203-0031330323032003-0031003101122201-0111122311033121-3111113123020202-1213020112221011) |
| `datadog_receiver.use_tls.enable_verify_certificate` | [datadog_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-2113130303131220-1003232101033221-3210011112112330-0210233330001323-2202312212223013-0033013320113303-3222031222102322-3110012032330021) |
| `datadog_receiver.use_tls.enable_verify_hostname` | [datadog_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-2211210201230223-3001032131120313-3333322310200213-3033130020200002-1021202101100223-1321200210130301-1302103021300031-2100021212121322) |
| `datadog_receiver.use_tls.mtls_disabled` | [datadog_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-1132320032323322-3320222303202233-1132013301000231-1131132021300223-2311113310023321-1011313203113212-3013231233022333-3331321201112021) |
| `datadog_receiver.use_tls.mtls_enable` | [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-2222101303102021-1330333221023011-3021301212011331-3101212213021203-3011213303333330-2203201203321312-3031023011212013-0002230013302210) |
| `datadog_receiver.use_tls.mtls_enable.certificate` | [datadog_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-0110312333011230-2112002111001111-1222131111302300-3320113102330100-2202230132100120-0101132310121201-1322230032212303-2323101023020012) |
| `datadog_receiver.use_tls.mtls_enable.key_url` | [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-2301132200033012-0310231230230211-0030103012130300-3322133320233213-3101132133220232-0333223211211311-0112012013122331-1213200322121112) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-3323322220333303-1200100303213031-0100030231211102-1020322013000131-1330201233032101-2311201203012032-2002120030130110-3011110003020320) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-1301121212310113-0000202230113010-0100303032311111-2030001213220222-1222203010203000-1221300101002233-2113311331120302-2320113103221013) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-2331203333033013-2323202032112313-3030111120131022-0332203131012001-0300321102001031-3131000223222231-0031303100031033-1113102213011322) |
| `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-3130331223111032-1121100132313231-2333302332011030-0031331310002211-0312103221110310-2021123110022223-3120010122010311-3232110021210220) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-2001133223133332-1130001323002101-0301011013332111-1322232223210310-2133221213333231-0013030333033013-0213130011303232-2101121303310230) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-002.md#canonical-1113320222230323-1112003133223110-0103033220221302-3232123323020202-2123231020311221-0120121301302120-3121111113232123-0211132120231030) |
| `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-002.md#canonical-2113222020302301-1032010030311033-2333302101310020-1201322122312321-2330200112211321-1201033211020121-2232031330303323-0331222113131312) |
| `datadog_receiver.use_tls.no_ca` | [datadog_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-002.md#canonical-2022202132121131-2222102202302010-2300230011332133-0302301131133300-3000223000323123-1102103203301233-1022101323122310-2110133113133101) |
| `datadog_receiver.use_tls.trusted_ca_url` | [datadog_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-002.md#canonical-3300132130321212-0032031020012231-1010030323300003-0131010102111020-0310313232132301-1022012200322313-1121232122322120-1001021321311211) |
| `description` | [description](data-sources--global_log_receiver--reference--group-001.md#canonical-1210131003031200-1302123200313011-3303320111231020-1030311102301232-2201122030001101-1230320233100002-0333010312223231-0233010330302113) |
| `dns_logs` | [dns_logs](data-sources--global_log_receiver--reference--group-002.md#canonical-0330003033103033-2232132203133203-1222002120211210-2331231322301220-1302122322201100-2332330003001322-1212032300002303-2110022032130313) |
| `gcp_bucket_receiver` | [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-0101022013031032-0020202300031330-1120302112000333-1303132333330203-3020123133110130-2002133033301033-0121030230213303-3033113231031210) |
| `gcp_bucket_receiver.batch` | [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-2312002313333013-0021201301213130-2323001123310120-3030331123100220-0202321312222103-0101000131231011-2130322202301230-2102133103222121) |
| `gcp_bucket_receiver.batch.max_bytes` | [gcp_bucket_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-002.md#canonical-2001223103202020-2032201133000230-2201101013112100-2311122030332100-0322131120301110-1210220032012322-0313220132032121-0112313011213330) |
| `gcp_bucket_receiver.batch.max_bytes_disabled` | [gcp_bucket_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-3231311101021003-0131011021303012-3120311310222110-0013012000103201-3000201333110121-0220112210333110-3122212122130213-3233312201002010) |
| `gcp_bucket_receiver.batch.max_events` | [gcp_bucket_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-002.md#canonical-1223210221011102-1321201133120302-2212022120012022-1231213221233013-3303323002033333-1113110303332301-1300131303221131-0331201203122322) |
| `gcp_bucket_receiver.batch.max_events_disabled` | [gcp_bucket_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-1001021320320233-3302113230123030-3233201232200001-1323102112100102-1230320311313211-3011102010110032-1302303320110222-0330231200203003) |
| `gcp_bucket_receiver.batch.timeout_seconds` | [gcp_bucket_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-002.md#canonical-3010323300330332-3110121103102301-1323312211010021-1003322112110302-3111211301303110-0201313020212032-2113232302311030-3030330133120003) |
| `gcp_bucket_receiver.batch.timeout_seconds_default` | [gcp_bucket_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-0202100310333002-1303301303213110-0330030123301010-2112113010203021-3033231102023333-3100010332311130-2213113110233132-2321223222231303) |
| `gcp_bucket_receiver.bucket` | [gcp_bucket_receiver.bucket](data-sources--global_log_receiver--reference--group-002.md#canonical-3012012110231221-3311321310103320-2303310200020310-3220223130120220-0132001223220011-3221323202223011-3330133220102021-1230332012102102) |
| `gcp_bucket_receiver.compression` | [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-3323333000331233-2320302013202111-1000121130313223-1011213303010133-0200020212012232-1300232133322013-1213202133212333-1220003100130011) |
| `gcp_bucket_receiver.compression.compression_default` | [gcp_bucket_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-2010133311333111-2100103031220313-3312010123330022-0221003033330332-3102301133232230-1311322303312112-1303213031223031-1203223120321033) |
| `gcp_bucket_receiver.compression.compression_gzip` | [gcp_bucket_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-0033111002032020-1331320011122111-2111101132311211-1000121121020301-2233311021003102-3310233123211003-1223212211000231-3012321331221033) |
| `gcp_bucket_receiver.compression.compression_none` | [gcp_bucket_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-1001113101011332-3331230122033022-3111211010132233-2320102021111111-3232222233021022-1213130133210132-2113002331221000-0120002302310330) |
| `gcp_bucket_receiver.filename_options` | [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-3310321102231203-3310202132031210-2121323231300210-0212210123331000-0321321022120002-0121321113120013-0313002011201212-0102201203003101) |
| `gcp_bucket_receiver.filename_options.custom_folder` | [gcp_bucket_receiver.filename_options.custom_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-1331332032221222-1111131202033310-2331321111210302-0001201322301322-1033110031323101-3121202021213211-0333101231300120-3230301123313112) |
| `gcp_bucket_receiver.filename_options.log_type_folder` | [gcp_bucket_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-2002012322213130-1122310113111002-0003232313322032-3203002320231020-3020220232200331-0221323001320301-2331121010221002-3230202032000231) |
| `gcp_bucket_receiver.filename_options.no_folder` | [gcp_bucket_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-1023312011011003-0111202123001222-3012020322130122-0120333200322313-3011123003321003-2103231001330033-0333220002000030-1213132300022121) |
| `gcp_bucket_receiver.gcp_cred` | [gcp_bucket_receiver.gcp_cred](data-sources--global_log_receiver--reference--group-002.md#canonical-2222200312302320-1222013032102102-1033313111212022-0100201013101200-1101122101201333-3133012021222021-3203221302333022-2232033003032221) |
| `gcp_bucket_receiver.gcp_cred.name` | [gcp_bucket_receiver.gcp_cred.name](data-sources--global_log_receiver--reference--group-002.md#canonical-0122221231323100-0200002202331002-3013120110200300-0331121001210220-0323330222211330-1220223323120233-2302002213331312-2202331020033103) |
| `gcp_bucket_receiver.gcp_cred.namespace` | [gcp_bucket_receiver.gcp_cred.namespace](data-sources--global_log_receiver--reference--group-002.md#canonical-1032201213232210-1213311003003223-2011101312101021-0212010232122023-2210221001033103-1223122322330132-2300010110302213-3132233223312232) |
| `gcp_bucket_receiver.gcp_cred.tenant` | [gcp_bucket_receiver.gcp_cred.tenant](data-sources--global_log_receiver--reference--group-002.md#canonical-0101101300330022-2233133122212233-3032110013220322-1311010232333311-1020222133202112-0312112212313201-1023022200332230-0012031222303121) |
| `http_receiver` | [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-0101022033231131-0331011310223210-1113000100233231-0202310200201130-3201032123020231-1122203030212220-0223101203210223-1100223030332223) |
| `http_receiver.auth_basic` | [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-1120223022332122-0033233120010022-1032231222332313-3321313103213103-0111123031201311-0223133220311332-3222011212123320-1221222102023232) |
| `http_receiver.auth_basic.password` | [http_receiver.auth_basic.password](data-sources--global_log_receiver--reference--group-002.md#canonical-3000010211001211-3230331311332221-3301313131300221-3120112001212131-0021033220312333-2121202222202231-2302131232220213-1212313203312001) |
| `http_receiver.auth_basic.password.blindfold_secret_info` | [http_receiver.auth_basic.password.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-0201302301101032-0302310111120020-0333231112312302-1033133130030201-0332233020113002-3123231100322202-0100112223332112-1233032031133011) |
| `http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-0330110232311231-3023100020333321-2201013032002213-0003330332303111-0323020230221010-1113030301011320-3101101102022222-0012302302303100) |
| `http_receiver.auth_basic.password.blindfold_secret_info.location` | [http_receiver.auth_basic.password.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-1100202133232122-1333132122230223-0110223322011223-0300201302130213-0212002113213200-1130331211323232-2133231222003123-2230123332120233) |
| `http_receiver.auth_basic.password.blindfold_secret_info.store_provider` | [http_receiver.auth_basic.password.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-3022322202113323-2202322323311321-2010323203022130-2320030032303313-0101101122211220-3200302311210220-0001022003011022-2100222111010302) |
| `http_receiver.auth_basic.password.clear_secret_info` | [http_receiver.auth_basic.password.clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-0013120022102030-3021321003030022-1132330013130103-1131211233311112-1120033103030020-3031010001130221-3221321103103000-2310211103010311) |
| `http_receiver.auth_basic.password.clear_secret_info.provider_ref` | [http_receiver.auth_basic.password.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-002.md#canonical-2313313001130133-1310000221230212-2212323222200201-2013211231023311-1131100233022210-3220011213101100-1220112011011232-1320311302103331) |
| `http_receiver.auth_basic.password.clear_secret_info.url` | [http_receiver.auth_basic.password.clear_secret_info.url](data-sources--global_log_receiver--reference--group-002.md#canonical-1222312013202132-1303300111321233-2031012213130203-0321211123131321-2300332102102323-3103000123332322-0011202111223300-3303002000320101) |
| `http_receiver.auth_basic.user_name` | [http_receiver.auth_basic.user_name](data-sources--global_log_receiver--reference--group-002.md#canonical-0330300300311121-1020331111221131-3003120230031300-0222311210320101-2113003211011132-2323032021223201-1132210023233233-1130213130221233) |
| `http_receiver.auth_none` | [http_receiver.auth_none](data-sources--global_log_receiver--reference--group-002.md#canonical-0001110103223111-1331200203113303-2333322331102123-3330222033232200-0002200002303333-2031300333230221-0120122223103023-1131311120220211) |
| `http_receiver.auth_token` | [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-002.md#canonical-2203000202112220-1133020201212101-0031233310023301-2031202223131123-2332210232322123-3320030020121211-2130301211222110-3102122202320212) |
| `http_receiver.auth_token.token` | [http_receiver.auth_token.token](data-sources--global_log_receiver--reference--group-002.md#canonical-0201021301103231-3001330030020313-0200101121102202-2030112032102122-2300331232201200-0312111003103310-3213211323101101-3130111001031131) |
| `http_receiver.auth_token.token.blindfold_secret_info` | [http_receiver.auth_token.token.blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-1012012301303002-1302322301202121-1000012132031113-1231221233301310-3222313323223112-1321233122103202-1203112213203100-3001321101121302) |
| `http_receiver.auth_token.token.blindfold_secret_info.decryption_provider` | [http_receiver.auth_token.token.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-1301032122200101-0101220020222202-2232121022203212-3311133210010123-2120313220103032-1213302200032232-1210312301312232-1210012101211312) |
| `http_receiver.auth_token.token.blindfold_secret_info.location` | [http_receiver.auth_token.token.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-002.md#canonical-0233123013332313-1102331232332101-3020200233000213-1101102002232032-2223212020310313-0320013201212121-0333021302210103-2012201003310322) |
| `http_receiver.auth_token.token.blindfold_secret_info.store_provider` | [http_receiver.auth_token.token.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-002.md#canonical-2020001013201213-1131120222230121-3312010002321331-1123020000321320-1312131332321203-1021213222232011-0230100110301013-0220131132200121) |
| `http_receiver.auth_token.token.clear_secret_info` | [http_receiver.auth_token.token.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-1101010302032113-1203123130203303-3220122113321002-2113031301001331-0123112023310021-2331231311001010-2031002312030112-3203101331230122) |
| `http_receiver.auth_token.token.clear_secret_info.provider_ref` | [http_receiver.auth_token.token.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-003.md#canonical-3133000002133303-2103223113130302-0111121112010311-0123131331213103-1223222001302020-1301132031013121-1210200123103121-0233222021201221) |
| `http_receiver.auth_token.token.clear_secret_info.url` | [http_receiver.auth_token.token.clear_secret_info.url](data-sources--global_log_receiver--reference--group-003.md#canonical-3333011102322313-3302133232110023-3031310310133201-1212320332310130-2100131311010201-0013010331102123-2232212331112123-3233001303332101) |
| `http_receiver.batch` | [http_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-3033002330211213-2112021121332210-0032223330233100-0232112000221301-0101211123332200-2111212201311020-1130122003111211-1201230201202100) |
| `http_receiver.batch.max_bytes` | [http_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-003.md#canonical-3120300312102212-2203000021131022-1120000010211023-1023012100011313-2230221201010213-0100320122111321-1010330211121020-0110200302030300) |
| `http_receiver.batch.max_bytes_disabled` | [http_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-0003000123021323-1203210121303101-1130112333130202-0212312121313220-0021001213033012-1131120323201311-2321202011321203-1323123301333111) |
| `http_receiver.batch.max_events` | [http_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-003.md#canonical-2221101300030012-3321002023010312-1322213010011003-3200311300100302-0020200203301110-2101310002232013-3000102321031103-1130230131220102) |
| `http_receiver.batch.max_events_disabled` | [http_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-1013132311021110-2211212013131133-0022002233233033-2303313010110012-1212313033112133-3010012102112030-0123311220033203-1310230101230000) |
| `http_receiver.batch.timeout_seconds` | [http_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-003.md#canonical-0102003133003300-1303311322102213-3212332232031303-0031013110013002-2213020320202111-1011213011001030-2100201113323131-0311032330122333) |
| `http_receiver.batch.timeout_seconds_default` | [http_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-0032213100301021-1200320130203012-1221102031021231-3223030013123122-2101232113013112-2232003223310313-2222031210301032-3100011331102013) |
| `http_receiver.compression` | [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1230110031111023-2110110002013300-2331231310011020-0030213032233130-0121121133031320-3301012011020222-1101230213201022-2332132321133123) |
| `http_receiver.compression.compression_default` | [http_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-1100121131310320-3232022323322030-2230333220200230-1311012303002011-2101201211100131-3112213313232330-3100103301211101-1222303030122322) |
| `http_receiver.compression.compression_gzip` | [http_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-2301200110202101-2112323302023131-0323300123313122-2230102222020032-3232323112132210-0022330123000220-3331332103013230-3100030110131230) |
| `http_receiver.compression.compression_none` | [http_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-3321022233022230-1131021301100122-0322000321023003-2330000103121000-2203233233200210-1132221011232232-2113123112021132-2210033101323203) |
| `http_receiver.no_tls` | [http_receiver.no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-1311130223100102-1002321010003212-3312331023231130-3133110201012032-0132132030023130-0222030110203332-2001333212132132-3203001321102320) |
| `http_receiver.uri` | [http_receiver.uri](data-sources--global_log_receiver--reference--group-002.md#canonical-2220110333311010-1121230323311311-3213322022222323-0120312332330312-3120302201133203-2131013211122023-2002130221002223-3122111322121030) |
| `http_receiver.use_tls` | [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2222103220232031-2031032010121130-2211221233312033-0332301331300312-2120012113132301-1213303302130011-2202233100003010-2022303330332120) |
| `http_receiver.use_tls.disable_verify_certificate` | [http_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-1012311323221212-1102020011022033-2323111232000032-1111101011332113-1331301002323202-0301331220010002-1003330322203033-1022303203233220) |
| `http_receiver.use_tls.disable_verify_hostname` | [http_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-2300022131232303-0333023333023330-3301102313132320-3003310311323320-0320312230101203-2003111011031200-1012120022112223-2021122200120100) |
| `http_receiver.use_tls.enable_verify_certificate` | [http_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-3232301210023212-0212212322110103-2200322003223232-1221013221010220-1121321120032311-2332122323030113-1331031302030120-2332110300022131) |
| `http_receiver.use_tls.enable_verify_hostname` | [http_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-3320320313031323-0100101101321230-0313122233131211-3313212130113120-3001211033122312-2100120310000110-3303320002133330-0011231032020230) |
| `http_receiver.use_tls.mtls_disabled` | [http_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-0010222301110230-0311211203023000-3010201033013031-0331230300332203-1213100321300013-0020200332123210-2113021203110132-3302233303222012) |
| `http_receiver.use_tls.mtls_enable` | [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-2113120302213213-1123011121122130-3212233323101200-3310302003320121-3223232330321201-0232202133330102-0021012322103320-2301311211223020) |
| `http_receiver.use_tls.mtls_enable.certificate` | [http_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-3011232331102121-2330330322002033-2201333131022212-3123023222300013-3123221233010012-3301202013213232-3202322230223333-3213002333212222) |
| `http_receiver.use_tls.mtls_enable.key_url` | [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-1113302222312031-0310021301203101-2021332030311232-1232130111100222-0303001200231112-0303012311213320-3202212012111113-3200112103330102) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-1313132100103300-2200020030302011-1223231310013233-0103203020020031-2003212223032200-3313032311120013-0321200000320303-1032002021021101) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-1321022223131123-1331320221102321-1122301032020222-3012132233001103-3132031312203023-1023303322300312-2300223303102230-0133331310112201) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-003.md#canonical-3003300203011300-3330100031310112-3123212021122332-3321031300113210-3031111322030303-0003300203233330-0133103031011231-2132201003100100) |
| `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-0021021202012031-2030011233031213-2020202332312033-2123101131202030-3130123233222133-2231320310030123-1132310301101033-2322001113221133) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-3213122221111213-0130101301002232-1023211111130201-2011230120121000-0013203010303301-0313131121311333-3133112230023002-1221331111133202) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-003.md#canonical-1012123123310222-3001133331320300-1001331222101302-1120312301001233-2212011231120332-3231312102211222-3023003001031130-0222320020333331) |
| `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-003.md#canonical-2021121112020000-3302012323231110-0002100122232233-1330322333103002-2323322132211202-2100123301003120-3012000313313032-3003000230002212) |
| `http_receiver.use_tls.no_ca` | [http_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-1130022333100312-2220220200023033-2002231330211320-3113330221112013-3110303002101020-2210033032213230-3231301230010201-1311100130023201) |
| `http_receiver.use_tls.trusted_ca_url` | [http_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-003.md#canonical-0302313312222112-0030303133011212-3021220013013321-3123013131232303-3211233023211032-2030212111022112-0133112030200120-0003100032123032) |
| `id` | [ID](data-sources--global_log_receiver--reference--group-001.md#canonical-2200111230003002-2201231231122133-1301001313221032-0012331210330233-1303000023223013-0223130312133233-0210312022013301-1120323033330321) |
| `kafka_receiver` | [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2203013201303101-3030330033021301-1103130222112033-1122003330112030-1222132131131010-1103203132001032-2212100031032010-3020032111211321) |
| `kafka_receiver.batch` | [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-1103322123332110-2312301222223013-3030323302031000-3033210312300312-0200330221132131-3012110022100232-2021001131233012-0002011100001301) |
| `kafka_receiver.batch.max_bytes` | [kafka_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-003.md#canonical-2012330020221010-1211133223123213-1200110031230300-2021131210323110-1032021203313023-1031201211223211-1111300301111030-3023013110313030) |
| `kafka_receiver.batch.max_bytes_disabled` | [kafka_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-3112302132301323-1021330023022202-0111232333331202-0300130322031002-1030022130133302-3033223132301201-2303113313220032-0203120131221221) |
| `kafka_receiver.batch.max_events` | [kafka_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-003.md#canonical-3303232321010230-0212022230010300-0000203323230112-1330122203020013-1132300202120332-3312122123322020-0131132302123210-3310102022322131) |
| `kafka_receiver.batch.max_events_disabled` | [kafka_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-3330200310000303-2103100211320103-1010331303313002-2202100201230130-3211200203200012-0031302130330332-1301010332101201-3220103200030020) |
| `kafka_receiver.batch.timeout_seconds` | [kafka_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-003.md#canonical-2311311020333333-0300321223100302-0023121312213012-1203021023302230-0030231012010132-0221311033221333-3333032311201332-3030000231333232) |
| `kafka_receiver.batch.timeout_seconds_default` | [kafka_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-3132213100123201-3103312231110023-1221200233103133-1222331202010101-1103021133320201-1300010213133301-0000212033130022-0032210302112303) |
| `kafka_receiver.bootstrap_servers` | [kafka_receiver.bootstrap_servers](data-sources--global_log_receiver--reference--group-003.md#canonical-1332023031001203-2013031021030231-2333220022313323-3231313313301031-1130012101013112-0211011313200030-2223223120313020-1032231122221331) |
| `kafka_receiver.compression` | [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-3122313310131001-0022220321013131-3310132233222032-2201132223011310-1233121232201212-1223222103112023-3302122222032323-1331330312133220) |
| `kafka_receiver.compression.compression_default` | [kafka_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-3020100312300133-1000020221310322-0333010222100321-1320231212103302-0101021311211113-1203021120112312-2322202303212202-0213010003200103) |
| `kafka_receiver.compression.compression_gzip` | [kafka_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-2011112302112310-3303121002333320-2220113131113232-3013210311202323-2210302001230003-0000022221121002-3120210123110001-2213000333120011) |
| `kafka_receiver.compression.compression_none` | [kafka_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-3201322230122102-2200203201202022-2001133323103333-0313122001013132-1313020333221231-1032123012310031-0211103130030101-2332311212020003) |
| `kafka_receiver.kafka_topic` | [kafka_receiver.kafka_topic](data-sources--global_log_receiver--reference--group-003.md#canonical-2012333120313110-1131020310021033-0032333230112121-0300112112131013-0122211121323213-0033323201201223-2003132232121221-0222322101330123) |
| `kafka_receiver.no_tls` | [kafka_receiver.no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2331001030323113-1111330223311321-3200012112103010-3333010100231012-3233301323323300-3201212022021030-1030010322201101-0212123032221022) |
| `kafka_receiver.use_tls` | [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3330132200100100-0020102213030203-2233233303310310-0302331011332221-2020111331103031-1012333222101213-0331200301100212-0232233220311331) |
| `kafka_receiver.use_tls.disable_verify_certificate` | [kafka_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-1223201221313231-0031123331012030-1030320120200222-2220312320121030-2333303213300310-3201211320023330-2231232033010012-2202302021202300) |
| `kafka_receiver.use_tls.disable_verify_hostname` | [kafka_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-1021002132200323-1222113013010021-0201231111300321-2000110302031311-2021130313333232-2220200301223021-0022310003211231-2320101022100033) |
| `kafka_receiver.use_tls.enable_verify_certificate` | [kafka_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-1103211111121003-0132101102312313-3220310300312323-3012120102112300-2032301000302131-0230022213221213-3333130301313132-2103211130222302) |
| `kafka_receiver.use_tls.enable_verify_hostname` | [kafka_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-3302011300132310-0001201322302032-0213032210102320-3212102222230210-3122233221232320-3033332223300020-2303331322332333-2201103203022210) |
| `kafka_receiver.use_tls.mtls_disabled` | [kafka_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-2012100003222003-1233022313301120-1202032301021033-2131330033301310-2300200000211030-3202023103132133-2003201201123232-1212300302130113) |
| `kafka_receiver.use_tls.mtls_enable` | [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-1033200222323211-2320212012032013-3131002201001211-2013131003222232-1203022133203000-2021031231321210-2112122300322010-2202002121230202) |
| `kafka_receiver.use_tls.mtls_enable.certificate` | [kafka_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-0122122313131013-1210211301200301-1211003300100120-0311211011232311-2203033120323220-0203201103000120-2132113123012311-0331233100002102) |
| `kafka_receiver.use_tls.mtls_enable.key_url` | [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-3311122312201101-1321301021320310-3220323330121130-1201303312223123-3300312323213101-0103033331232022-1233001321313330-2231101120111233) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-0011011202120123-0021102131000112-3210332212001312-1031121022013120-0312332002221022-0023130221310331-0002300203011102-3022013223202011) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-1222112120321313-2313031003233123-2300131113101213-1212022030112223-0023221110313133-3102301311112221-1221110010113320-1332302131101211) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-003.md#canonical-2210003301020133-0333331332133030-3031331300323302-1133230210022031-1121032202001020-1020123322020101-1302132031103333-2300321211133130) |
| `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-2303221022012133-0310013031323222-2312023323111101-3203131313123231-1333332000302110-2232012130030033-1220123123323230-0232121022031010) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-2201011011210212-0111313312010303-3300222233031030-2130331000323300-3033202002200033-1303022313230333-1120330003022233-2001311103233300) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-003.md#canonical-2121112010001300-2222021231100003-0132200232330133-1111121122100011-3300303310332212-2200332121003012-1312120020310230-2332012133031121) |
| `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-003.md#canonical-2030131233002212-2120212302120310-3323113020220031-2102233313023301-1120311233101131-3103123301010231-0323321323102022-3123211313311203) |
| `kafka_receiver.use_tls.no_ca` | [kafka_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-2203302101301230-1301203331100223-1032001221101212-1103110311322233-3032330010321003-3030113030302222-2203130222002333-3121011132121102) |
| `kafka_receiver.use_tls.trusted_ca_url` | [kafka_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-003.md#canonical-0231203313002022-1201122033032121-2323203200101030-0013002101233001-2233330320211102-2002300300301022-0230301231132210-1303300201021310) |
| `labels` | [labels](data-sources--global_log_receiver--reference--group-001.md#canonical-2110212111112102-1220100112002313-3122230022332000-0313223100220331-1100133322321120-2321120232332320-2231113132032002-3123011133021003) |
| `name` | [name](data-sources--global_log_receiver--reference--group-001.md#canonical-0000131312002213-1233112220201021-2002310231232330-2230202330211020-3222221223213313-3101121233223132-0110013201123103-3201000102330023) |
| `namespace` | [namespace](data-sources--global_log_receiver--reference--group-001.md#canonical-1011201023030032-2202132102310112-1121133112301202-1203123302021102-1330333111131013-0233010133013200-0020201331003133-3200230012123121) |
| `new_relic_receiver` | [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-3132102203031223-0300022203231020-2033232013233211-3131300201011022-0210221130300123-2213111030003231-1002103002302030-1301031211011001) |
| `new_relic_receiver.api_key` | [new_relic_receiver.api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-1300211121201200-2001033220100202-1220231130130320-0012300201031003-0332101013010033-0213323210133310-2311223210302030-2231013130131301) |
| `new_relic_receiver.api_key.blindfold_secret_info` | [new_relic_receiver.api_key.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-1231103320103313-0102121031223030-2020310333300022-2323333332210101-0002332301310230-0123212010130013-1211123003310003-1132310233032003) |
| `new_relic_receiver.api_key.blindfold_secret_info.decryption_provider` | [new_relic_receiver.api_key.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-3132200230032303-0233301133220223-3032032223020331-2023113011302323-0012322003311013-2222230022123322-1101213120232013-1220321010031022) |
| `new_relic_receiver.api_key.blindfold_secret_info.location` | [new_relic_receiver.api_key.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-003.md#canonical-0023020303322230-0331201222021023-1130131230003033-3003300320320332-1302012331211332-1101001013122122-1303121302022012-2210332312003021) |
| `new_relic_receiver.api_key.blindfold_secret_info.store_provider` | [new_relic_receiver.api_key.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-003.md#canonical-2303332202301320-2010103010213131-1311021030310322-0302330302103011-0302313123130101-0112023013123310-3110122123330021-3211210013102210) |
| `new_relic_receiver.api_key.clear_secret_info` | [new_relic_receiver.api_key.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-2323323320223303-3313101021211333-2301212210131022-2201131020203033-3112310311333132-3213130123202303-2110301011331223-1312030212011001) |
| `new_relic_receiver.api_key.clear_secret_info.provider_ref` | [new_relic_receiver.api_key.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-003.md#canonical-0021210323230210-0122010012121103-2132130022322213-2231300003033233-0201110023113311-3123033000133323-2023013122330331-2232103200220021) |
| `new_relic_receiver.api_key.clear_secret_info.url` | [new_relic_receiver.api_key.clear_secret_info.url](data-sources--global_log_receiver--reference--group-003.md#canonical-1211122201032320-0213133231011120-3332212030010222-3030013103103303-0231033203102122-3231023133011302-3311231202321332-0333230323212212) |
| `new_relic_receiver.eu` | [new_relic_receiver.eu](data-sources--global_log_receiver--reference--group-003.md#canonical-2221130203203000-3332120110030003-2022323311313023-0133113023001023-1211033233002101-3201012312233330-0301012221233111-1210121212320222) |
| `new_relic_receiver.us` | [new_relic_receiver.us](data-sources--global_log_receiver--reference--group-003.md#canonical-3302121330210213-0222031021210020-0021100300123132-2333202101020311-2103330133110230-2021332010033231-2132332113232213-2200020021210010) |
| `ns_all` | [ns_all](data-sources--global_log_receiver--reference--group-003.md#canonical-1030120112010331-2230012313100203-3132133213000333-1232233330012101-3233322300200100-0011332212003301-2110203211000302-1101000123003013) |
| `ns_current` | [ns_current](data-sources--global_log_receiver--reference--group-003.md#canonical-2200320110300010-3121021332001011-2113033313031130-3212321003321130-2223103133200223-0313021133310311-2200213300131303-3022101333300002) |
| `ns_list` | [ns_list](data-sources--global_log_receiver--reference--group-003.md#canonical-3011000200300333-3300013030233023-1333110132103131-0000012030212002-1311112102002200-3110313322002301-2002302330131301-1220022231112010) |
| `ns_list.namespaces` | [ns_list.namespaces](data-sources--global_log_receiver--reference--group-003.md#canonical-2313233010230233-3102101333201310-0211211210103002-2031201102321311-2323001203320232-0300122212203320-2222220003322323-3130132302111012) |
| `qradar_receiver` | [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-1203320123322112-0003220233010101-0221021023303230-1311333320232133-2123212302200211-3330303321331111-3000310130102210-0323203022331310) |
| `qradar_receiver.batch` | [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-0120311201302311-3111210120211230-2122322000322312-3121321102322130-0201003100023323-1031310221123110-1113000112023303-1112322133203001) |
| `qradar_receiver.batch.max_bytes` | [qradar_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-003.md#canonical-3320011112101313-0001012302000200-2131121221330311-1333213212103130-1301023010220101-0001212321120222-3020201322000230-0213003113230003) |
| `qradar_receiver.batch.max_bytes_disabled` | [qradar_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-3202322303333020-3303230033331023-1201223030333100-3012213120120320-2200000201122112-1032223133233331-0123331313110312-0010320310011321) |
| `qradar_receiver.batch.max_events` | [qradar_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-003.md#canonical-0113221213010220-3233322030121102-1210122003310100-1231231221201231-0003210210030030-2313311101322011-3221211230120202-1113020023013122) |
| `qradar_receiver.batch.max_events_disabled` | [qradar_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-3101000111023132-2300323013100103-2131022300011113-2032323020123303-1100332032023321-1103303321201100-3112200011110133-0300302313213231) |
| `qradar_receiver.batch.timeout_seconds` | [qradar_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-003.md#canonical-1101001232300231-3202223232301000-2322123000330201-0333302321032133-1010322103121022-2021031201332232-2322210001131320-2100321320111120) |
| `qradar_receiver.batch.timeout_seconds_default` | [qradar_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-2013111300311330-1132223122303330-3211202100302101-3302133113020101-2220213212112020-2223201233000102-3320020330002220-0101103133132313) |
| `qradar_receiver.compression` | [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-0303032320311320-0122021323231033-0230220210031231-1100133131230133-2332121013033202-2023113010300300-0000121330111301-0032023111332002) |
| `qradar_receiver.compression.compression_default` | [qradar_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-1033113110202130-1121122202002322-3010013321332001-1323131133032231-3312203100210202-2113303332112311-1102322112221223-0001213121222003) |
| `qradar_receiver.compression.compression_gzip` | [qradar_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-3233321202021213-1230302303101303-1311033122322023-0111302223230222-3033131101333112-0023100202311013-3031312221330122-0003132100302210) |
| `qradar_receiver.compression.compression_none` | [qradar_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-1030203010111302-0132031201202133-2130302302311311-1023013202200013-2031312021213332-3023310332121022-1201022101003101-3112332223303023) |
| `qradar_receiver.no_tls` | [qradar_receiver.no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-1311221010330131-0100030122201013-0123222200333120-2221022230023332-0133123010211132-0300332331311303-2303010213100302-1300031202112222) |
| `qradar_receiver.uri` | [qradar_receiver.uri](data-sources--global_log_receiver--reference--group-003.md#canonical-1312101322230223-1133303032220313-2202202302223232-2100022232122321-2210100103132030-3120331221320220-2000210101030023-0023130122110010) |
| `qradar_receiver.use_tls` | [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2230220303112210-1332303332023323-0303300310322223-1021010001233001-2201000221131021-3113101123232032-1210221030233301-0323022013221302) |
| `qradar_receiver.use_tls.disable_verify_certificate` | [qradar_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-1010110110313311-3022111300213103-2021321201103031-0010220303330001-3120330212103221-3311222332331301-2313001223210133-3102112211323120) |
| `qradar_receiver.use_tls.disable_verify_hostname` | [qradar_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-0033032130120233-2223021003101311-3101310203222221-2100201013200230-2312010311131120-1121120223111102-1210313300121101-0222021301301201) |
| `qradar_receiver.use_tls.enable_verify_certificate` | [qradar_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-3103312302111301-1231132032330333-3021000133010132-1332232332313323-1211033212323012-0210102013132320-3201021000322220-3212331133103003) |
| `qradar_receiver.use_tls.enable_verify_hostname` | [qradar_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-2332003033202112-1303203011320320-2331011201020301-0322103001220033-0302010232213302-1223032102123333-0003313100001303-2303001121133003) |
| `qradar_receiver.use_tls.mtls_disabled` | [qradar_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1031002111110322-0020113230113332-0323211303022220-0203030213301032-0003032131113013-0300110032210130-2320310122133121-1223213301132030) |
| `qradar_receiver.use_tls.mtls_enable` | [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-1102000323013231-3011332031221221-3223322211322130-3133331323223121-3232303103233103-0331221222131112-1223101302213111-3221301213001220) |
| `qradar_receiver.use_tls.mtls_enable.certificate` | [qradar_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-1023232113213120-0000322100312001-2123103300323311-1302120103233201-0201210131220111-2232202012112303-2122203032030310-1003310133231322) |
| `qradar_receiver.use_tls.mtls_enable.key_url` | [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3013222211230321-0300321101212300-0131201130103201-0110310111322300-1322020301322001-0002300130110132-1012112222220011-1002132110023013) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-0033312112311302-2211223000112211-3020201001322220-0313222322010323-2310213203232110-0301332033031013-0211220200102303-0310133302001210) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-0322330321221331-1013132123001032-2101210332311030-0013320031333000-2032313230201312-0030310231233321-1331013030230103-3302233212110333) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-004.md#canonical-3221232030221232-3013021202033023-2322132321313123-2200002001232131-0301103100332021-3022330111022110-0100131320022201-0002231122112022) |
| `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-1202122332200302-2222131032210313-1003230301231330-0123303132113322-2122113331120122-2311012020003120-2221001130132223-3032000321201012) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-2201210312331202-0101132022131022-2222010023212102-2023013013203102-3312023322011122-3033102011133110-2031032213123303-3132011310320302) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-004.md#canonical-0133230231030101-2220032310201011-3301300030132332-3121222122111013-2002203001121223-0100022131033210-1112000031321221-1230202330333311) |
| `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3133211121033330-0233013321011211-0020022001231222-3330023303233130-0322030301001130-2223303333312030-1110033220130211-2000002233032102) |
| `qradar_receiver.use_tls.no_ca` | [qradar_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-2022312312103013-1230211221100032-2332102120220202-2023002223022230-3202111112303012-1120332111111212-0022023231311011-0001333322320123) |
| `qradar_receiver.use_tls.trusted_ca_url` | [qradar_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-003.md#canonical-3323001230003112-2032112220213030-2223112202201301-2022033202030013-3103202312020201-0202121133200210-1023023220031012-3100220201130302) |
| `request_logs` | [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1032022001202310-0203211112112032-3223133020101320-1122331023331223-0011221222002323-0101233233111212-1002013321121320-1323300223110231) |
| `request_logs.sampled` | [request_logs.sampled](data-sources--global_log_receiver--reference--group-004.md#canonical-0213000030003023-1302322002320330-1101232212233301-1333011030312222-0232332313312220-3002233121133303-1131323201303233-0201300101121120) |
| `request_logs.unsampled` | [request_logs.unsampled](data-sources--global_log_receiver--reference--group-004.md#canonical-2302132121301121-0112133011310222-1221312213031332-3202011003220203-1221000010321133-3012123231311323-1213012233300221-1012211031332330) |
| `s3_receiver` | [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3212111221313232-1220032220002121-1013323021001210-2102032312000332-1100102021232231-2013023201232020-2333203332302002-2022012110300003) |
| `s3_receiver.aws_cred` | [s3_receiver.aws_cred](data-sources--global_log_receiver--reference--group-004.md#canonical-3202012023230310-1232321333123111-2110301020021221-1222030102300112-1022302333122211-0002023120303122-3310223101003302-1232311223201111) |
| `s3_receiver.aws_cred.name` | [s3_receiver.aws_cred.name](data-sources--global_log_receiver--reference--group-004.md#canonical-0322202223031112-3032032230302130-2330322222210210-2212330001231202-0213232332031201-0130003313021110-1121222111233103-0212002031023222) |
| `s3_receiver.aws_cred.namespace` | [s3_receiver.aws_cred.namespace](data-sources--global_log_receiver--reference--group-004.md#canonical-1300113000103321-2023233201122030-3113320102221202-2113021312022100-2011012330213333-3132130100213212-2010022121233111-2321133221202022) |
| `s3_receiver.aws_cred.tenant` | [s3_receiver.aws_cred.tenant](data-sources--global_log_receiver--reference--group-004.md#canonical-1313220200133133-0303012021322010-2322333212132001-0113210103003000-0102000301112103-3231322330303333-2231133012130130-0103231211302020) |
| `s3_receiver.aws_region` | [s3_receiver.aws_region](data-sources--global_log_receiver--reference--group-004.md#canonical-3130112022220123-0322010213101030-2030001220321011-3222113132131011-3021021310310021-2112332203022012-2013312231100122-0030030133210210) |
| `s3_receiver.batch` | [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-3020230212103033-2133200232032131-3030102231220220-0212110022322010-1113213223103133-1201013033203131-3323110321302031-2023031103332131) |
| `s3_receiver.batch.max_bytes` | [s3_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-004.md#canonical-1232003320120121-1030132103012101-3300312103030100-3012333211001311-1020133010333002-2030223013330000-3302300302123300-1231222033030221) |
| `s3_receiver.batch.max_bytes_disabled` | [s3_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1031312331231133-2030212110221311-2021111101203223-2110103313222202-0013031132001221-1011313230222130-1223031002230130-2312311031102020) |
| `s3_receiver.batch.max_events` | [s3_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-004.md#canonical-1221000313212123-3110311020011320-1131332220203311-0011331333033001-1102111323231121-0112300121002313-1121202210323000-1222032103221003) |
| `s3_receiver.batch.max_events_disabled` | [s3_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-3123332003223121-2213310010110002-1023023102211100-2232123322210112-2133113023311001-2021021032101303-0001130231320313-1213021231002320) |
| `s3_receiver.batch.timeout_seconds` | [s3_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-004.md#canonical-1200121311230210-3122321333020103-0120222203010002-3232310321210101-0120232232013331-1121233123022031-1013302232323233-3231203010011013) |
| `s3_receiver.batch.timeout_seconds_default` | [s3_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-1223311312003032-2320022133231021-1221010003222300-1223300001213203-3311231021211012-0232320301211112-0310102303102010-3220030131330313) |
| `s3_receiver.bucket` | [s3_receiver.bucket](data-sources--global_log_receiver--reference--group-004.md#canonical-0010213232000030-1120300322233310-2021103020303311-3220221331323312-0210232131011112-3122323121311010-1032300103222012-2032010201003301) |
| `s3_receiver.compression` | [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3001333213310201-2102302233230011-3032230023101000-3232110002022301-1102001231322133-2030103332321002-3002320222013201-0312013121300213) |
| `s3_receiver.compression.compression_default` | [s3_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-2101011132302010-1030032330223332-2331120003000330-0021322010212233-2103201031201120-1030331312213122-1321222100100313-2202013222120010) |
| `s3_receiver.compression.compression_gzip` | [s3_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-1220000312113133-2210122210013121-3002312323100031-2332022331021101-1020122102103021-0322231130303011-2012323132101103-0020110213121231) |
| `s3_receiver.compression.compression_none` | [s3_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-0222021302001020-1002323122223230-2121022131132310-2000023221200211-1311330112010322-3011000012030021-3011202100113013-0200222130110312) |
| `s3_receiver.filename_options` | [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-1032230031213311-2300200122322002-0302230112310111-2120100032200020-2301110213210003-2021332303201111-1322333322030322-3113032031023012) |
| `s3_receiver.filename_options.custom_folder` | [s3_receiver.filename_options.custom_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-3230322132020033-0122030010132232-1220200103000203-3133103023310103-0333033321222301-3212011323001322-3230000102021202-0103220122222322) |
| `s3_receiver.filename_options.log_type_folder` | [s3_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-2322233312020231-2000120120233333-3113002301321110-0000132302223232-3132211033310322-3232132330332321-2202312231021022-3322112001213221) |
| `s3_receiver.filename_options.no_folder` | [s3_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-3211122120332313-3121320303230030-0222131322213202-2002322003202101-3222031120323101-1030111030312000-2110032103222302-1132220333031010) |
| `security_events` | [security_events](data-sources--global_log_receiver--reference--group-004.md#canonical-2121201001133031-1031120231001131-1132311112122102-3321300232323210-0223013202122312-3113023102021320-3231333333203113-1212202033021010) |
| `splunk_receiver` | [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2130322203222220-0110323213010302-1233212331033202-0013332333331213-2311200220221030-0313332130213320-3232333202032213-1331302023310301) |
| `splunk_receiver.batch` | [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0310220231331303-0331002120123031-0301111231212212-1023313103123122-0000330132030210-3020131210312123-1200013032123101-0001321123123300) |
| `splunk_receiver.batch.max_bytes` | [splunk_receiver.batch.max_bytes](data-sources--global_log_receiver--reference--group-004.md#canonical-1313123330123213-0003011032101233-3100023111111030-1132130103203023-0023311121132023-0211032310032013-2223222211222223-0103022301200332) |
| `splunk_receiver.batch.max_bytes_disabled` | [splunk_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1232023303133101-3020111132101123-0102313223233313-2300012130123310-2122102302312132-2312110110030130-0023122000311322-2103232313120211) |
| `splunk_receiver.batch.max_events` | [splunk_receiver.batch.max_events](data-sources--global_log_receiver--reference--group-004.md#canonical-0332001220230311-1010200201130011-2030132303020020-3211320112201213-1023222330200313-2110312100010333-2322322023310323-2132301201311211) |
| `splunk_receiver.batch.max_events_disabled` | [splunk_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1311200213112200-3231012032111330-2021101211203231-1033333300110112-0230312323032022-1223132213220201-3133212203021103-0200330313121323) |
| `splunk_receiver.batch.timeout_seconds` | [splunk_receiver.batch.timeout_seconds](data-sources--global_log_receiver--reference--group-004.md#canonical-0210302133331323-2002310321031133-2310301132233113-2223020302300102-2000213122103310-1100222200330101-3210300131221100-3123021020202220) |
| `splunk_receiver.batch.timeout_seconds_default` | [splunk_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-0102310221311100-3321300201103310-1120201321320102-2200031123021023-2313311331312130-2013200120021233-0020310121011101-3213322313303103) |
| `splunk_receiver.compression` | [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1132323220133001-0201123332100322-1033310330132112-3213213013322122-0300012121231123-3020321013321232-3323213333203302-1030032221133032) |
| `splunk_receiver.compression.compression_default` | [splunk_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-3311331031122302-0212312022022233-0312032020232210-2132203310211031-0103311001021322-1002222320000023-3211103312123213-0221312003323302) |
| `splunk_receiver.compression.compression_gzip` | [splunk_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-2132130303021302-2101200321303130-0221321111211022-2233003020031223-3323311320110311-1021220110220012-1303032031122203-3211100221123100) |
| `splunk_receiver.compression.compression_none` | [splunk_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-2221100001033010-3102322311003213-2110202223011002-3112322101332001-0301201130313111-1322300310132002-2120010331112000-2012303132000132) |
| `splunk_receiver.endpoint` | [splunk_receiver.endpoint](data-sources--global_log_receiver--reference--group-004.md#canonical-0122310212212110-3112223030311123-2232032230233300-3030032021332100-0032011231230113-2103000003021332-0101003333210312-0013311330230221) |
| `splunk_receiver.no_tls` | [splunk_receiver.no_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-1323320233332023-2333133131013212-0300231313001332-3002202330132300-0231200131032222-3223322211102210-2011000310220322-1312000232220302) |
| `splunk_receiver.splunk_hec_token` | [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-0010103202313023-3011132320331202-0102212312100212-3001200030210211-2311320200030020-3322033230203133-3023103200120212-2003331010321021) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info` | [splunk_receiver.splunk_hec_token.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3132210202011312-0201101331120113-1222232030200210-1230121333300323-1230101102232033-0020302320112213-2303000333023102-3032303211133120) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-2303101132032203-0211112123202323-3211110030200203-0131223211012130-1110332210123011-2332310301332103-0210313012320213-1111132032323103) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.location` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-004.md#canonical-0103031121311131-3323313022123223-3203111111300213-2300231302300202-0232120303332032-1101033203231301-1103202232102110-3013020000113203) |
| `splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider` | [splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-1311131020313032-1133313011023101-3023031222233011-0100300020213101-1311132232113303-2121100212213012-1122100101202033-3133123220013330) |
| `splunk_receiver.splunk_hec_token.clear_secret_info` | [splunk_receiver.splunk_hec_token.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-0200023233103100-3011023230000130-0100301222313331-3121101032301320-0231302002301322-3101002221230110-1032013322310103-2122320232021113) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref` | [splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-004.md#canonical-1333313132331203-1003130333221001-3222233022303012-3223123031320111-2321100003122123-1122303121202133-0033232202131203-2101121321213203) |
| `splunk_receiver.splunk_hec_token.clear_secret_info.url` | [splunk_receiver.splunk_hec_token.clear_secret_info.url](data-sources--global_log_receiver--reference--group-004.md#canonical-0231130312323030-1130312230200320-2320111130021202-2131123210301103-1320111220031212-0201110010111301-2122203012333333-3321312303021333) |
| `splunk_receiver.use_tls` | [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-1031212113011132-2011333230211332-2123311222333320-3331202123213021-1113210222203122-3231203212213130-1301320311201323-1313312332103102) |
| `splunk_receiver.use_tls.disable_verify_certificate` | [splunk_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-2223231332002313-0032022223310312-1231223133230100-2123121200320213-3033232222021103-2322032320322231-1123100133131031-3003321022111031) |
| `splunk_receiver.use_tls.disable_verify_hostname` | [splunk_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-2220022321101202-3023033101013113-3113332220001112-2111312011023122-2301330200013003-3113232311121320-3130132320023333-2002100320323032) |
| `splunk_receiver.use_tls.enable_verify_certificate` | [splunk_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-0230322230221202-2020112013000230-3123100122122031-1133022021313122-2031120331313201-1320223023301333-1021003222302010-2012313213103321) |
| `splunk_receiver.use_tls.enable_verify_hostname` | [splunk_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-2310123001312230-0021301000230220-1033231002320021-2222230121103303-0033022312001232-2131331000033020-3300210221033122-2323000113022022) |
| `splunk_receiver.use_tls.mtls_disabled` | [splunk_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-3012003232022201-3102013121131022-0232111011323021-2000213212120120-3133132033333333-0323100303331113-1110320011033330-1103000320110013) |
| `splunk_receiver.use_tls.mtls_enable` | [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-2010302101030012-1210121202101231-1111311133202220-2021132103100332-3012001222103222-1023303233013021-2101111032233310-0101130022223233) |
| `splunk_receiver.use_tls.mtls_enable.certificate` | [splunk_receiver.use_tls.mtls_enable.certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-0120010122112010-0002132022130120-2032300031113301-3100221130213203-2120213321103212-0323302233311113-3330022223033103-2032233232232220) |
| `splunk_receiver.use_tls.mtls_enable.key_url` | [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3113201231201020-0013002203113121-0333123202230231-3113112321232032-1132112323111003-1211202112312232-3121010132010333-3012103323301131) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-2120321122323330-0221123321123322-0012200100202233-0213210301302221-0120131120112130-0133201332300111-2003321133012323-3233130130300122) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-3102322113210211-3303301020102000-0201200131101202-1120120113220313-1032223012312300-3130033323031032-2321220201232212-0110200311032301) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-004.md#canonical-0322110130022031-1030130230320232-2101300030000133-1323310321002003-2233010231323330-3003321030223122-1230303112011132-0003030330130130) |
| `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` | [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-1321131122030021-0010103220020013-0003002011330200-2320030102300313-3002321222230111-1333001301231323-1323130202223303-0011323200102022) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3232103321112020-0023111312201010-0222130313012211-2213300313122012-0100033123321213-0233220100131002-3321211203112113-0003320232222233) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-004.md#canonical-0301133101332311-1330020030200330-2300331102002000-1021120321000120-1001022003301003-1330300323101302-2121313320110321-3121220131213222) |
| `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` | [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3012023233110300-2100323311133121-2320013322321211-0321233302300023-3323203232303231-1103221001020023-2200121332021022-2233331333210333) |
| `splunk_receiver.use_tls.no_ca` | [splunk_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-0121120030121223-3311221203133322-0331001300031213-2123200101302232-2330321212013011-1212222223213010-0333131113101330-2311122221133123) |
| `splunk_receiver.use_tls.trusted_ca_url` | [splunk_receiver.use_tls.trusted_ca_url](data-sources--global_log_receiver--reference--group-004.md#canonical-1033200111011313-3230202023223222-2101030323123302-1003323303012031-2123100222300223-0111100333233002-2231311332002310-0220231120112202) |
| `sumo_logic_receiver` | [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2111201203021003-1010212310123022-2023221323331032-2212330321202303-1123222120131130-2113333103311120-3110301220010121-0200302122002201) |
| `sumo_logic_receiver.url` | [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3031113231133321-2013020123113033-3211311322001231-0323202201000203-0230303230001011-1131213213200200-0200233011132202-2110303323102013) |
| `sumo_logic_receiver.url.blindfold_secret_info` | [sumo_logic_receiver.url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-2323021002101033-1133310322210301-3120301223233331-3200112230302133-2031000301200123-0302300201133113-1211233300003302-2312300100313001) |
| `sumo_logic_receiver.url.blindfold_secret_info.decryption_provider` | [sumo_logic_receiver.url.blindfold_secret_info.decryption_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-1121232332123313-0120222303303113-2311222222302330-2122130020010333-1302110313023030-2122021101322213-3103010122102113-0110210101310032) |
| `sumo_logic_receiver.url.blindfold_secret_info.location` | [sumo_logic_receiver.url.blindfold_secret_info.location](data-sources--global_log_receiver--reference--group-004.md#canonical-2300330231310021-1302012222330123-1310131131321132-2111021102122021-0020030123013003-1133032123120100-0102301331101023-1112121223002011) |
| `sumo_logic_receiver.url.blindfold_secret_info.store_provider` | [sumo_logic_receiver.url.blindfold_secret_info.store_provider](data-sources--global_log_receiver--reference--group-004.md#canonical-0301223313113202-1011312230223030-2022203332200310-1023102133313112-1120213221223033-2202331111013232-3322221021130132-3013323030231013) |
| `sumo_logic_receiver.url.clear_secret_info` | [sumo_logic_receiver.url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-0021203223313310-2211132223321003-3203301313223033-2031021001330323-0123202023020213-0021323022133130-1023113021001123-1313121020213002) |
| `sumo_logic_receiver.url.clear_secret_info.provider_ref` | [sumo_logic_receiver.url.clear_secret_info.provider_ref](data-sources--global_log_receiver--reference--group-004.md#canonical-1303300110310232-2312131212230133-3211332002003123-1012111001131130-3201302200320122-3231113033320302-0320333210202031-3311112222222012) |
| `sumo_logic_receiver.url.clear_secret_info.url` | [sumo_logic_receiver.url.clear_secret_info.url](data-sources--global_log_receiver--reference--group-004.md#canonical-1213131123010212-2301033211021332-0222321211113120-0202330202030231-0320200010301032-1102121133202132-2110231032122033-1320213100220203) |

<a id="canonical-2102122231133210-0022102033021322-3003211213232101-3131031031022021-1030300013113012-3311103201011132-3013200312021233-1212021132311102"></a>

## Next pages — Property reference / 220322330022 / 11

- [audit_logs](data-sources--global_log_receiver--reference--group-001.md#canonical-3133212020032302-0323212203031203-3200200231033010-2303311030030012-3001102312131320-3233231303203033-1101020321312133-2201321333233323)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-0132210313021033-2221133303001132-2002201122230200-2203313221112330-3122130230230331-3110223323220223-1131013312112132-1123200222220321)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-2121301220330030-1221312033213003-0320321302002303-2133212312111123-3330122101022123-0133210113001032-2201030003210013-3211030023023312)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [dns_logs](data-sources--global_log_receiver--reference--group-002.md#canonical-2313012101200221-3103103122211233-0010031133112232-3010013110321112-0132120212211113-2313113031330310-0333130212030033-3300322103021000)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-1222203022011122-1213111132203013-1201312120332200-3233301331202013-3012133313300230-3320033231321122-3201112101101230-3232031212020130)
- [ns_all](data-sources--global_log_receiver--reference--group-003.md#canonical-3222301230023030-0001112112122320-3100220123111200-3201120222111311-1330032121023233-1201222312320311-1111331322201210-0133312301312020)
- [ns_current](data-sources--global_log_receiver--reference--group-003.md#canonical-3213312020222012-2002220201233002-0031302031132022-2311120330010033-0031032202323202-0022102301032121-0321103001233122-0001120123201310)
- [ns_list](data-sources--global_log_receiver--reference--group-003.md#canonical-3303120033132301-0202221321231331-0130123312332131-0202000200222232-3021322301011113-1000013103130021-2133210101223030-0302330311022122)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [security_events](data-sources--global_log_receiver--reference--group-004.md#canonical-1002031033031310-1233001100003213-2003102123102332-1202233033221123-1021111002311202-0121010220310110-1211132100130223-2001223323310023)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3133212020032302-0323212203031203-3200200231033010-2303311030030012-3001102312131320-3233231303203033-1101020321312133-2201321333233323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321010201031113-0133320021030000-0101022100100010-3032312110200133-0221120011313302-2032132112231101-0111130113130313-1112021003320300"></a>

## audit_logs — audit_logs / 012313021002 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- audit_logs

<a id="canonical-2133221203001223-3123231200122131-3110232122221212-2033103202301212-0132123303320010-3111333302033122-1110213131202301-2111222131331002"></a>

Type: `["object", {}]`. Computed.

\[OneOf: audit\_logs, DNS\_logs, request\_logs, security\_events\] Enable this option

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

OneOf alternatives in this subsection:

- [audit_logs](data-sources--global_log_receiver--reference--group-001.md#canonical-2133221203001223-3123231200122131-3110232122221212-2033103202301212-0132123303320010-3111333302033122-1110213131202301-2111222131331002)
- [dns_logs](data-sources--global_log_receiver--reference--group-002.md#canonical-0330003033103033-2232132203133203-1222002120211210-2331231322301220-1302122322201100-2332330003001322-1212032300002303-2110022032130313)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1032022001202310-0203211112112032-3223133020101320-1122331023331223-0011221222002323-0101233233111212-1002013321121320-1323300223110231)
- [security_events](data-sources--global_log_receiver--reference--group-004.md#canonical-2121201001133031-1031120231001131-1132311112122102-3321300232323210-0223013202122312-3113023102021320-3231333333203113-1212202033021010)

Select alternatives according to the provider validators above.

<a id="canonical-3211212111112210-3220311312320311-3312330323010132-0331110312312233-2313110302332211-0111331313133301-1331003322020333-3122133203330121"></a>

## Direct properties — audit_logs / 012313021002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333221331000313-1332020100011010-3100133022220221-2223203120320000-0233211022132011-0311020123202332-3321201232130121-1231220231303133"></a>

## Next pages — audit_logs / 012313021002 / 4

- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002233001102300-2003032030132022-3202210123013102-2111013002120033-0121223221011213-1220002010213010-3320323030233022-3322330011033221"></a>

## aws_cloud_watch_receiver — aws_cloud_watch_receiver / 023300300330 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- aws_cloud_watch_receiver

<a id="canonical-3331110112131012-1000100132111012-1203111230222102-1230112233112113-0001200333200231-0312301323230303-1011311201201322-2301312200332113"></a>

Type: `"single"`. Computed.

\[OneOf: aws\_cloud\_watch\_receiver, Azure\_event\_hubs\_receiver, Azure\_receiver,
datadog\_receiver, gcp\_bucket\_receiver, http\_receiver, kafka\_receiver, new\_relic\_receiver,
qradar\_receiver, s3\_receiver, splunk\_receiver, sumo\_logic\_receiver\] AWS Cloudwatch Logs
Configuration for Global Log Receiver.

Upstream description:

AWS Cloudwatch Logs Configuration for Global Log Receiver.

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

- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-3331110112131012-1000100132111012-1203111230222102-1230112233112113-0001200333200231-0312301323230303-1011311201201322-2301312200332113)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1133203103010030-0213101233221030-0113301022130000-0112220222113332-2113012101223231-1210310130311121-1322212023012203-3120213302333232)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-0303310011011220-0200131310311211-2300302321112220-1311210110210130-2100302022132032-3210212112313131-3213323010012223-2203111201011331)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1030122021022232-1220312302320123-0120123303110130-0320113101000121-3332333032132102-1213013023033133-1200133200111110-3230023131300203)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-0101022013031032-0020202300031330-1120302112000333-1303132333330203-3020123133110130-2002133033301033-0121030230213303-3033113231031210)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-0101022033231131-0331011310223210-1113000100233231-0202310200201130-3201032123020231-1122203030212220-0223101203210223-1100223030332223)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2203013201303101-3030330033021301-1103130222112033-1122003330112030-1222132131131010-1103203132001032-2212100031032010-3020032111211321)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-3132102203031223-0300022203231020-2033232013233211-3131300201011022-0210221130300123-2213111030003231-1002103002302030-1301031211011001)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-1203320123322112-0003220233010101-0221021023303230-1311333320232133-2123212302200211-3330303321331111-3000310130102210-0323203022331310)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3212111221313232-1220032220002121-1013323021001210-2102032312000332-1100102021232231-2013023201232020-2333203332302002-2022012110300003)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2130322203222220-0110323213010302-1233212331033202-0013332333331213-2311200220221030-0313332130213320-3232333202032213-1331302023310301)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2111201203021003-1010212310123022-2023221323331032-2212330321202303-1123222120131130-2113333103311120-3110301220010121-0200302122002201)

Select alternatives according to the provider validators above.

<a id="canonical-1221222110330101-1111020000321202-0110300132222323-2300223120022322-0231023033222020-1220000320231320-0311201013122112-3111211211310032"></a>

## Direct properties — aws_cloud_watch_receiver / 023300300330 / 3

- [aws_cred](data-sources--global_log_receiver--reference--group-001.md#canonical-0331111001222030-0221210122022212-2133313103313010-0123221213101313-0312302312110310-1201333112020001-0212301001123031-3002313011113330): complete subsection reference.

<a id="canonical-3311102200012103-3032303221112130-0113102011002001-2133021101120132-1110010032220322-3312010332100033-1320120110331131-1233212302121121"></a>

<a id="canonical-2213310101121030-0021331302332132-1103320012223132-3020201001112222-2100201000002210-1222002210210120-3133321100113223-3120322013013030"></a>

## aws_region property — aws_cloud_watch_receiver / 023300300330 / 4

Type: `"string"`. Computed.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Upstream description:

AWS Region Name.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [batch](data-sources--global_log_receiver--reference--group-001.md#canonical-3200223321012003-3232102021300202-0101332211223132-0222310012212121-2133300120012123-1320023023321312-1333100312132133-3103321213101320): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-001.md#canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113): complete subsection reference.

<a id="canonical-0013332233001311-3323103011023333-2313200223201223-3013332031232222-3330333211033331-0220120010222303-1303222023012103-0100221120021302"></a>

<a id="canonical-0322232013020010-1313303132000131-1010303211231230-2323100230212221-2131013123211021-3033112101201333-2232212223000221-0303012100232200"></a>

## group_name property — aws_cloud_watch_receiver / 023300300330 / 5

Type: `"string"`. Computed.

The group name of the target Cloudwatch Logs stream.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  }
}
```

<a id="canonical-1022333110331123-1031010001031013-0332110031122222-1002112102020010-1330311002201211-1301311213222120-1032013031023220-0303332030021303"></a>

<a id="canonical-3133022011311220-1101001121202321-1130223032213230-0331231103001032-3112222030122331-3133331001333123-2113323203031101-3323100311330110"></a>

## stream_name property — aws_cloud_watch_receiver / 023300300330 / 6

Type: `"string"`. Computed.

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

Upstream description:

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[^:*]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[^:*]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[^:*]*$"
  }
}
```

<a id="canonical-0332113110001202-2213303313302103-2033120223023222-0311022102000012-1130030202321302-0332112330212203-2311122210211022-1213213000321221"></a>

## Next pages — aws_cloud_watch_receiver / 023300300330 / 7

- [aws_cloud_watch_receiver.aws_cred](data-sources--global_log_receiver--reference--group-001.md#canonical-0331111001222030-0221210122022212-2133313103313010-0123221213101313-0312302312110310-1201333112020001-0212301001123031-3002313011113330)
- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-3200223321012003-3232102021300202-0101332211223132-0222310012212121-2133300120012123-1320023023321312-1333100312132133-3103321213101320)
- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0331111001222030-0221210122022212-2133313103313010-0123221213101313-0312302312110310-1201333112020001-0212301001123031-3002313011113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023113102201012-3032003331303110-2330012111133230-3321302033232202-2211301321303323-2111112032212011-2021022023003311-0020320132231003"></a>

## aws_cloud_watch_receiver.aws_cred — aws_cred / 110112023322 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- aws_cloud_watch_receiver.aws_cred

<a id="canonical-3121101011123321-1230031310012031-3011202211032002-0233302301311202-1232122303112311-2232210033101011-0301200003030332-1300023211110232"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3100011301100130-1310221221220201-1122221101113213-3321010310133303-3312313033122233-0221302211012232-0211102213313020-3201121202231222"></a>

## Direct properties — aws_cred / 110112023322 / 3

<a id="canonical-0202130221020102-1002120220201021-2101131131032010-3300122130203203-2333101121121223-1332323121320200-0020000012333230-3233013202320221"></a>

<a id="canonical-1011312032312303-2100021320003012-2110023200313320-0023130232320322-3023000321100320-1021031230032222-3312002312120222-2330232233022201"></a>

## name property — aws_cred / 110112023322 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1103211111032210-1321331101303301-1200323312010012-0303013330012030-1002222032321301-1101001013112323-0132232132321121-3110203220300033"></a>

<a id="canonical-2322211313301203-0130031100311323-0331332110322011-0133103321211323-0200332212333203-0300213303212121-2200221010012330-3312311313333233"></a>

## namespace property — aws_cred / 110112023322 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3003003123111310-3301010313313211-1012230001230313-2031032002322300-3132030211313020-2121001111202112-3132333113210113-0322331330320201"></a>

<a id="canonical-1013132133223223-0213300222222211-1303102022031101-0000012012002331-3023210030331033-0013031022311032-1230001301111203-3123212302323203"></a>

## tenant property — aws_cred / 110112023322 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2201332213012012-3333220022300203-2221331230123022-1103310202021320-1133130223032100-0333122031002330-2012220331231001-0231310031122102"></a>

## Next pages — aws_cred / 110112023322 / 7

- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3200223321012003-3232102021300202-0101332211223132-0222310012212121-2133300120012123-1320023023321312-1333100312132133-3103321213101320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012100020123023-0102221133121113-2102001333220123-1321202220113331-1311030110211133-3030020201302011-2203123102010201-1123221300321133"></a>

## aws_cloud_watch_receiver.batch — batch / 311120311102 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- aws_cloud_watch_receiver.batch

<a id="canonical-1233331313232013-0210102231321101-2233132230112030-1102102102322100-0210030231312322-2131002102132002-2121002032120031-0113012201312130"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-3323201210221000-1320300123202000-1223000233021033-3102230010302130-3302031030230210-3300233300200123-1210202232001131-0301210231303211"></a>

## Direct properties — batch / 311120311102 / 3

<a id="canonical-3111222320110330-0211233113013301-0130332330113333-0332103102301221-3113032313132223-0003223213322131-3333031300120331-0133313030310133"></a>

<a id="canonical-3230122233320231-3312332320333001-1302110303211123-1203033020202013-2221323220002303-0011313231330303-3310213130011223-3212112231212021"></a>

## max_bytes property — batch / 311120311102 / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-0233133132202111-2300210002003230-3210202000003311-2222112300033022-2122300200033002-3230003011021022-3302203123331233-2200233231002010): complete subsection reference.

<a id="canonical-2130002032131222-2211200213033111-2232010310021320-1310010021110130-1232311111232010-2311131202213023-3002132101110331-3000211222203220"></a>

<a id="canonical-2200002011322202-1232313130233032-0103100020132020-1121322122233311-2200113131131012-1020220301332003-1012233020111211-0211331200323013"></a>

## max_events property — batch / 311120311102 / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-0011102000021103-2003103210221302-2233102001301032-0200211301112322-2221120330100232-2332021331300311-3310213320330001-3310033321031100): complete subsection reference.

<a id="canonical-0020211302321221-2113321213220001-1030001002023023-1223112032202120-1033230003020322-2000022013013333-3121121103110132-2113023121100310"></a>

<a id="canonical-2021233022110221-2021123123212212-1302333031223231-3130220230131133-3121001000033133-0303231021212302-0032013203320300-2230022002201002"></a>

## timeout_seconds property — batch / 311120311102 / 6

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-001.md#canonical-3003032330323213-1203233233322101-0121003102110300-0121001132213002-1202302203323133-0100300120302320-1220231111213220-2032023232030033): complete subsection reference.

<a id="canonical-0031211320222132-2010121101200110-3210023122123003-3322002122031110-1302122031323223-0100200330212122-2111023031300223-0311121312000310"></a>

## Next pages — batch / 311120311102 / 7

- [aws_cloud_watch_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-0233133132202111-2300210002003230-3210202000003311-2222112300033022-2122300200033002-3230003011021022-3302203123331233-2200233231002010)
- [aws_cloud_watch_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-0011102000021103-2003103210221302-2233102001301032-0200211301112322-2221120330100232-2332021331300311-3310213320330001-3310033321031100)
- [aws_cloud_watch_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-001.md#canonical-3003032330323213-1203233233322101-0121003102110300-0121001132213002-1202302203323133-0100300120302320-1220231111213220-2032023232030033)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0233133132202111-2300210002003230-3210202000003311-2222112300033022-2122300200033002-3230003011021022-3302203123331233-2200233231002010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301231103332102-0322211322123103-2330003313102230-1011020011302223-2213023223111232-1212332111033212-2213022322322023-0000310200223122"></a>

## aws_cloud_watch_receiver.batch.max_bytes_disabled — max_bytes_disabled / 032132100133 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-3200223321012003-3232102021300202-0101332211223132-0222310012212121-2133300120012123-1320023023321312-1333100312132133-3103321213101320)
- aws_cloud_watch_receiver.batch.max_bytes_disabled

<a id="canonical-1310033223231011-2332311222312102-1321322133332200-3333101123123000-1210330113013203-1233031223120101-1010210133222321-1120130101322321"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0223220300203112-1021201112012130-2322301120032200-2032133212030102-3020121101223220-1011200013100311-0031232213303011-3131111231320232"></a>

## Direct properties — max_bytes_disabled / 032132100133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303322132011312-2010031130301201-2301311113003103-2023331022010111-3332111123211100-2132311121200133-3211022100130022-3200222202012120"></a>

## Next pages — max_bytes_disabled / 032132100133 / 4

- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-3200223321012003-3232102021300202-0101332211223132-0222310012212121-2133300120012123-1320023023321312-1333100312132133-3103321213101320)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0011102000021103-2003103210221302-2233102001301032-0200211301112322-2221120330100232-2332021331300311-3310213320330001-3310033321031100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113022331122230-3323101332132330-3023311010313212-0220121323301002-1013031310302120-3011200310213303-3111032201121300-3221123223302013"></a>

## aws_cloud_watch_receiver.batch.max_events_disabled — max_events_disabled / 222212303103 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-3200223321012003-3232102021300202-0101332211223132-0222310012212121-2133300120012123-1320023023321312-1333100312132133-3103321213101320)
- aws_cloud_watch_receiver.batch.max_events_disabled

<a id="canonical-2302110232231000-2323222320201120-1202001230202313-3033132222223122-1012331302331033-0230002113201223-0333132033313100-2113020031213232"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2132031020101111-2231203223131002-0111301211232033-2232103333022310-3121233303002103-1220121013122002-2202013122220133-3112103230101302"></a>

## Direct properties — max_events_disabled / 222212303103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002021201102210-1313003202031222-0012202330233313-2232211000223002-3220203201030211-1103010321202313-3230322212133130-3103102021123132"></a>

## Next pages — max_events_disabled / 222212303103 / 4

- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-3200223321012003-3232102021300202-0101332211223132-0222310012212121-2133300120012123-1320023023321312-1333100312132133-3103321213101320)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3003032330323213-1203233233322101-0121003102110300-0121001132213002-1202302203323133-0100300120302320-1220231111213220-2032023232030033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302120310233121-2310202123312021-2320230313300130-1210020223233022-2112010333333111-0330212232100232-2333100232330031-3022203213112322"></a>

## aws_cloud_watch_receiver.batch.timeout_seconds_default — timeout_seconds_default / 331301003022 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-3200223321012003-3232102021300202-0101332211223132-0222310012212121-2133300120012123-1320023023321312-1333100312132133-3103321213101320)
- aws_cloud_watch_receiver.batch.timeout_seconds_default

<a id="canonical-2312120211222333-2032311303102131-3000122003213310-2023032013321331-3131121321113101-0131112102220231-2202003011330311-2031132130122100"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0101232102130011-2310222321313121-1010313223013330-1110233110100201-3122021023330210-3320221101323003-0101213020313131-0212311103232303"></a>

## Direct properties — timeout_seconds_default / 331301003022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321023120303331-1231210321110131-0201211002011102-1300301323032302-2230002201231020-3303032132120233-3013013310123033-3022001220232230"></a>

## Next pages — timeout_seconds_default / 331301003022 / 4

- [aws_cloud_watch_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-3200223321012003-3232102021300202-0101332211223132-0222310012212121-2133300120012123-1320023023321312-1333100312132133-3103321213101320)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212010303330320-3111022113310322-3212031211020201-0323220130330313-2310330310333312-0303002123133122-1113333331132323-2321003310321311"></a>

## aws_cloud_watch_receiver.compression — compression / 200223212031 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- aws_cloud_watch_receiver.compression

<a id="canonical-2110121120030013-2001320103312120-2103313323000223-0321121211112313-0101200213310000-2031202201132121-0223212033011110-0332311102312013"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-0203320212221103-1231312002121111-1033112021210220-0302013122112012-1133012011313133-2012200313113123-1033003123311113-0130220320210310"></a>

## Direct properties — compression / 200223212031 / 3

- [compression_default](data-sources--global_log_receiver--reference--group-001.md#canonical-2133000003212110-3320102202310130-2100322333121003-1032133003101230-0211130030223311-1201012311011132-3213203330231330-0003033332113002): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-001.md#canonical-1233130202031121-0331011012001223-3221010100231030-2313231233210200-2320233030233223-3100101032311100-3303022312333002-3130020132232001): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-001.md#canonical-3000031310000101-0320013113031101-1130102302003333-3031321332100202-3221231012122300-3033113021011013-1131023031231220-1330333013223110): complete subsection reference.

<a id="canonical-2312232303002010-0132310023120013-2113223132011020-3122120120010203-0211333323201000-2321123123002220-1103101131232111-1333031320122301"></a>

## Next pages — compression / 200223212031 / 4

- [aws_cloud_watch_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-001.md#canonical-2133000003212110-3320102202310130-2100322333121003-1032133003101230-0211130030223311-1201012311011132-3213203330231330-0003033332113002)
- [aws_cloud_watch_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-001.md#canonical-1233130202031121-0331011012001223-3221010100231030-2313231233210200-2320233030233223-3100101032311100-3303022312333002-3130020132232001)
- [aws_cloud_watch_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-001.md#canonical-3000031310000101-0320013113031101-1130102302003333-3031321332100202-3221231012122300-3033113021011013-1131023031231220-1330333013223110)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2133000003212110-3320102202310130-2100322333121003-1032133003101230-0211130030223311-1201012311011132-3213203330231330-0003033332113002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012023210123312-2202133233313203-3201021223031032-1022210320131101-1213211310310010-1023122320330220-1203002231012222-0221122333202012"></a>

## aws_cloud_watch_receiver.compression.compression_default — compression_default / 313102120132 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113)
- aws_cloud_watch_receiver.compression.compression_default

<a id="canonical-1030032322103222-1133311133112000-2303221023222200-2023221031233320-2331300230013000-3113200330130132-2212202131211123-3103123212133302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-0111311200311032-3300212023203020-3302230013133101-3113122130211122-3232203122000123-2133031111002100-2320232223123030-3001013221110020"></a>

## Direct properties — compression_default / 313102120132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111103111230123-2103232313202000-3303023001020333-2033023120123011-3111232100110113-2023130222331112-3121220030222302-2233102330013201"></a>

## Next pages — compression_default / 313102120132 / 4

- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1233130202031121-0331011012001223-3221010100231030-2313231233210200-2320233030233223-3100101032311100-3303022312333002-3130020132232001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022212203003021-0030130113303132-3323111301202331-1322000333030230-2322203013222313-3031223123022010-1313020201201120-3310300132333031"></a>

## aws_cloud_watch_receiver.compression.compression_gzip — compression_gzip / 202100312130 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113)
- aws_cloud_watch_receiver.compression.compression_gzip

<a id="canonical-1223223221203132-0331322110213000-3202231011223233-3001001102323313-1213111220211133-3101313220122203-3233101230133023-3212122011020030"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3210300221302103-1020202022011103-0110302002102103-3302033003300123-3213300031300212-0002230002023313-1021122213111100-3312310231013330"></a>

## Direct properties — compression_gzip / 202100312130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122100322130333-1223123313200211-3312213221212301-1022002211122203-2032210301311102-1130233123220333-0323222332220300-2311200332033323"></a>

## Next pages — compression_gzip / 202100312130 / 4

- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3000031310000101-0320013113031101-1130102302003333-3031321332100202-3221231012122300-3033113021011013-1131023031231220-1330333013223110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031020201001331-0020011213331302-0111122001302000-1021321020110230-2010332001323103-2323023311121033-0103100312032201-1113123212132013"></a>

## aws_cloud_watch_receiver.compression.compression_none — compression_none / 221313220101 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [aws_cloud_watch_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1211213330230033-0032123010133001-0201012100331222-1012011301010333-1201022022301110-2302103220201210-2100333230232120-3110032202312032)
- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113)
- aws_cloud_watch_receiver.compression.compression_none

<a id="canonical-0012320110221120-2013320300301201-0011031230103110-2030221131121102-1332023310233231-3131321313331031-1202102231233210-1203032313302332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-1020231112330333-2322012323211213-3031033330232010-0021311220120123-3233132300001222-2101122201133031-2013001333131313-3222322013032022"></a>

## Direct properties — compression_none / 221313220101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222130100321101-1310132102210031-3003233022102302-3113110121313102-3302300323202122-2112310023101303-1303020031132003-2010213201221012"></a>

## Next pages — compression_none / 221313220101 / 4

- [aws_cloud_watch_receiver.compression](data-sources--global_log_receiver--reference--group-001.md#canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0132210313021033-2221133303001132-2002201122230200-2203313221112330-3122130230230331-3110223323220223-1131013312112132-1123200222220321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011100112230001-0123311320302300-1331331112312331-3321303003000032-2203310323132003-3120222022310011-2332201213202300-2032101313101222"></a>

## azure_event_hubs_receiver — azure_event_hubs_receiver / 031012001301 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- azure_event_hubs_receiver

<a id="canonical-1133203103010030-0213101233221030-0113301022130000-0112220222113332-2113012101223231-1210310130311121-1322212023012203-3120213302333232"></a>

Type: `"single"`. Computed.

Azure Event Hubs Configuration for Global Log Receiver.

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

<a id="canonical-0112300300032230-2111210323122300-0323122312222111-0222103233203032-2110103103010011-1311201222230233-0212223113110222-3303131231312010"></a>

## Direct properties — azure_event_hubs_receiver / 031012001301 / 3

- [connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-2223231000011202-1020213011311200-3022103320232230-1031201322122213-0102311120232012-0212311232030203-2021110113013021-2213021303231003): complete subsection reference.

<a id="canonical-3233003101121100-2112231100120032-1012002220321100-0120103120021231-2301023112312020-2333231213300302-0211111022020323-2033221033211200"></a>

<a id="canonical-1133331010333112-1010033210313032-2133031321122332-0233011301203300-0210020320223230-3201301012101033-2133130201121120-3330130312113320"></a>

## instance property — azure_event_hubs_receiver / 031012001301 / 4

Type: `"string"`. Computed.

Event Hubs Instance name into which logs should be stored.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
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
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

<a id="canonical-3212012100330311-0001100130100130-3323321211212101-0311033223300313-0033220010120001-0223010303230200-3012221313212321-1222232211330022"></a>

<a id="canonical-2203033030333322-0302331222203323-2220221101301132-2033020000312001-0103230233002313-2001132330310312-1230020310210321-0300303100001113"></a>

## namespace property — azure_event_hubs_receiver / 031012001301 / 5

Type: `"string"`. Computed.

Event Hubs Namespace is namespace with instance into which logs should be stored.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$",
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

<a id="canonical-1102012323311003-1300200321201202-1033121323120100-3120331033123010-0003001121213101-2021032331330221-2333232300221210-3302212132133220"></a>

## Next pages — azure_event_hubs_receiver / 031012001301 / 6

- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-2223231000011202-1020213011311200-3022103320232230-1031201322122213-0102311120232012-0212311232030203-2021110113013021-2213021303231003)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2223231000011202-1020213011311200-3022103320232230-1031201322122213-0102311120232012-0212311232030203-2021110113013021-2213021303231003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303232321300233-0213323022103210-1010321001100312-3323223230220311-0203001333011021-3101121013003221-3221030032020313-1132030202102323"></a>

## azure_event_hubs_receiver.connection_string — connection_string / 031100320212 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-0132210313021033-2221133303001132-2002201122230200-2203313221112330-3122130230230331-3110223323220223-1131013312112132-1123200222220321)
- azure_event_hubs_receiver.connection_string

<a id="canonical-0333100212131103-1133011001333101-2212011112023001-3200022001301210-1200012031211302-2001013303132330-1333301001202001-3101030311232032"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-3311011133231011-0130322200100320-2131032012130013-2220032131000020-0301312121233303-0032303222133312-0130030123132012-1300033233202023"></a>

## Direct properties — connection_string / 031100320212 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-1023330331121031-3212230311231113-1222000320121132-3020003133233203-0121123033311313-2310213001022122-3110133123021302-0331210000201213): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-2221321220021310-3331003020030122-1312123100033002-2330032221331021-3303210111122203-2010002232110331-0113303030133300-2023000022211231): complete subsection reference.

<a id="canonical-2131321000212112-3111302102103332-1002033210123202-2322032201032032-3101302031303230-2121301211233323-0222231212201020-3332130110300112"></a>

## Next pages — connection_string / 031100320212 / 4

- [azure_event_hubs_receiver.connection_string.blindfold_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-1023330331121031-3212230311231113-1222000320121132-3020003133233203-0121123033311313-2310213001022122-3110133123021302-0331210000201213)
- [azure_event_hubs_receiver.connection_string.clear_secret_info](data-sources--global_log_receiver--reference--group-001.md#canonical-2221321220021310-3331003020030122-1312123100033002-2330032221331021-3303210111122203-2010002232110331-0113303030133300-2023000022211231)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-0132210313021033-2221133303001132-2002201122230200-2203313221112330-3122130230230331-3110223323220223-1131013312112132-1123200222220321)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1023330331121031-3212230311231113-1222000320121132-3020003133233203-0121123033311313-2310213001022122-3110133123021302-0331210000201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303102320331031-0333000312023110-1111112201301223-3012031202122131-3212310303103130-2032302121020233-3131202013122303-0321312020011232"></a>

## azure_event_hubs_receiver.connection_string.blindfold_secret_info — blindfold_secret_info / 123000300310 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-0132210313021033-2221133303001132-2002201122230200-2203313221112330-3122130230230331-3110223323220223-1131013312112132-1123200222220321)
- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-2223231000011202-1020213011311200-3022103320232230-1031201322122213-0102311120232012-0212311232030203-2021110113013021-2213021303231003)
- azure_event_hubs_receiver.connection_string.blindfold_secret_info

<a id="canonical-1233121333012322-0212112222222323-1330121130103201-0232112020111102-0220213113221220-0100102010110321-1100302120020220-3101131233210330"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-2222202202333213-3021220032202010-3001121313201003-3012123011230201-0201213310210331-2023203110223313-0133220130221123-3100032211011131"></a>

## Direct properties — blindfold_secret_info / 123000300310 / 3

<a id="canonical-2131033323211322-0200320002201333-3020130303303032-1332212123120111-2310032020200013-0112121130023211-2223210120322301-3020101111030133"></a>

<a id="canonical-3102232231123120-2223111103222000-3222033131130200-2012033132313120-1030131312213012-3222123210302031-3210302031011233-2330030023100110"></a>

## decryption_provider property — blindfold_secret_info / 123000300310 / 4

Type: `"string"`. Computed.

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

<a id="canonical-2221123102131101-3003030122112302-1322023112122132-2020330332121133-1332112033222232-2231200322310331-1100231322230011-0131103023113132"></a>

<a id="canonical-2212313321231113-1003131130133123-1011001210222101-0103221100022200-3223102223232022-2113010121133302-1133023303130030-0133210110300312"></a>

## location property — blindfold_secret_info / 123000300310 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0202201213021030-1313123303303022-3130220112012123-1302032232101313-3223231113110010-2200311112031022-1312022002221300-3233022311102233"></a>

<a id="canonical-1130232323121001-0121031120033102-1232123231220111-1332021301230102-3212331023121003-2013110311011002-0102201220202001-3323332211120320"></a>

## store_provider property — blindfold_secret_info / 123000300310 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-1002213201212121-1031202210320032-3133020230330200-1122122131011000-0032111122230131-1111232001211310-3300200011002133-2201203213310300"></a>

## Next pages — blindfold_secret_info / 123000300310 / 7

- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-2223231000011202-1020213011311200-3022103320232230-1031201322122213-0102311120232012-0212311232030203-2021110113013021-2213021303231003)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2221321220021310-3331003020030122-1312123100033002-2330032221331021-3303210111122203-2010002232110331-0113303030133300-2023000022211231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303013333201321-2103313222303202-2131000323312223-2223000121323322-0133230332102112-2202203013333313-0300311121133330-2120001002121301"></a>

## azure_event_hubs_receiver.connection_string.clear_secret_info — clear_secret_info / 303010123312 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [azure_event_hubs_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-0132210313021033-2221133303001132-2002201122230200-2203313221112330-3122130230230331-3110223323220223-1131013312112132-1123200222220321)
- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-2223231000011202-1020213011311200-3022103320232230-1031201322122213-0102311120232012-0212311232030203-2021110113013021-2213021303231003)
- azure_event_hubs_receiver.connection_string.clear_secret_info

<a id="canonical-1130233320221032-1321331300233201-1322133322211221-1333221021321010-0210232333033320-1313200022332122-3222320230232103-0012030210003332"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-2230000011200101-1023332301230110-3033301223233312-3221120120302033-1130100223300032-0100223001333323-1000201112210033-1211131310220103"></a>

## Direct properties — clear_secret_info / 303010123312 / 3

<a id="canonical-3331033000310013-2231323032112021-2013021210110332-3002131323323202-0312311013032012-1220303122100231-1110101012333022-0102003132333110"></a>

<a id="canonical-1213203112200310-3020210120220222-1102010200131122-0131223011133130-2332211112211210-3331000020101213-0013110100322111-3031221120300233"></a>

## provider_ref property — clear_secret_info / 303010123312 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2103133030323122-3301103203000210-2101212302202333-0233033123231121-2101122311233323-0033010231203302-2302301133223322-0322203202201213"></a>

<a id="canonical-3111023312120120-1321311032331021-2302131000212100-3123112100213103-3203203101012131-0011022231322112-3313211022123213-0010223003133220"></a>

## URL property — clear_secret_info / 303010123312 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0210323000033230-2110023220102123-2332112320110122-1320232010120030-3333013212000200-1333021032203301-1112132013223030-1032210220010203"></a>

## Next pages — clear_secret_info / 303010123312 / 6

- [azure_event_hubs_receiver.connection_string](data-sources--global_log_receiver--reference--group-001.md#canonical-2223231000011202-1020213011311200-3022103320232230-1031201322122213-0102311120232012-0212311232030203-2021110113013021-2213021303231003)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2121301220330030-1221312033213003-0320321302002303-2133212312111123-3330122101022123-0133210113001032-2201030003210013-3211030023023312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113130212031313-1200222232321230-3223221113021332-2213023030331022-2010130300300202-2203200133122113-0201213010300120-1200033103311003"></a>

## azure_receiver — azure_receiver / 311003033102 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- azure_receiver

<a id="canonical-0303310011011220-0200131310311211-2300302321112220-1311210110210130-2100302022132032-3210212112313131-3213323010012223-2203111201011331"></a>

Type: `"single"`. Computed.

Azure Blob Configuration for Global Log Receiver.

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

<a id="canonical-0103310012213022-3302122200333113-0330321203113310-3220111311121102-0232023311022223-3102123321013211-1130213222230001-2002021110301113"></a>

## Direct properties — azure_receiver / 311003033102 / 3

- [batch](data-sources--global_log_receiver--reference--group-001.md#canonical-2110231131123022-3210021331120313-0230130302332321-2122330301322100-0032000123010111-3000120033100202-2313222223213221-0231002202020032): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-002.md#canonical-1301222013203232-0131031232320300-0333013231220113-1222032002300113-2232001032022102-0233013102310320-0100023112330001-0320231331113332): complete subsection reference.

- [connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-2301311003132033-3012223003123133-1213311332201031-2031121222131030-3011202200333201-1112312313312323-1202123022021223-1113232332011210): complete subsection reference.

<a id="canonical-1122122002230032-1001101212113032-0210102101211011-2133220013310021-2211211020032000-2012313301021132-1313321133323322-2002023220030010"></a>

<a id="canonical-0220330111223022-0102001011323000-0101033210303212-1303301010322330-0223011113233103-1110202220320302-3321312011131320-1230030313111331"></a>

## container_name property — azure_receiver / 311003033102 / 4

Type: `"string"`. Computed.

Container Name is the name of the container into which logs should be stored.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 3,
    "pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
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
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
  }
}
```

- [filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-1213320220313111-2133213301031032-3323200320132221-3331102313013310-3021210320112221-3123112202113322-0220322200312200-2222020203030033): complete subsection reference.

<a id="canonical-1123221012022122-3132111001033331-2031333010102011-2313212320022132-2302002000133111-2233231031230301-0112332000310102-2323221230331232"></a>

## Next pages — azure_receiver / 311003033102 / 5

- [azure_receiver.batch](data-sources--global_log_receiver--reference--group-001.md#canonical-2110231131123022-3210021331120313-0230130302332321-2122330301322100-0032000123010111-3000120033100202-2313222223213221-0231002202020032)
- [azure_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-1301222013203232-0131031232320300-0333013231220113-1222032002300113-2232001032022102-0233013102310320-0100023112330001-0320231331113332)
- [azure_receiver.connection_string](data-sources--global_log_receiver--reference--group-002.md#canonical-2301311003132033-3012223003123133-1213311332201031-2031121222131030-3011202200333201-1112312313312323-1202123022021223-1113232332011210)
- [azure_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-1213320220313111-2133213301031032-3323200320132221-3331102313013310-3021210320112221-3123112202113322-0220322200312200-2222020203030033)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2110231131123022-3210021331120313-0230130302332321-2122330301322100-0032000123010111-3000120033100202-2313222223213221-0231002202020032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211220333120300-3333232201012001-2302210003323200-1130313110233020-1230132331010313-1202010313301032-3210111111122132-0302233132203130"></a>

## azure_receiver.batch — batch / 130300120132 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-2121301220330030-1221312033213003-0320321302002303-2133212312111123-3330122101022123-0133210113001032-2201030003210013-3211030023023312)
- azure_receiver.batch

<a id="canonical-2323221120203333-2211030001033223-0010121333110011-3201032021133322-3130332033033302-2322031012132313-1202021230002111-2223313330110213"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-0002221313312013-1301300032033011-2012203023303023-2322323302322311-2000033001030200-1033221113303000-2030221101031100-0000112022223323"></a>

## Direct properties — batch / 130300120132 / 3

<a id="canonical-2322030301221021-2103122000031332-3322331030122111-2023200231111133-2330213331211012-0233023012023012-2221322302000233-0120120123001312"></a>

<a id="canonical-2023000212300023-2331110323331120-1032001100031110-0122311121212132-2332333030211332-1333311001212200-3320231032102331-2330320213221330"></a>

## max_bytes property — batch / 130300120132 / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-0303111300332213-1202311101313321-2032103010002210-1001010131030121-3220023313310212-0032232112021333-1232211101312003-3132000132033303): complete subsection reference.

<a id="canonical-2220011103021002-0300030232032122-1112001113300003-3131102331313031-3303111312122130-2020223222013032-2030200122010033-0300031220132212"></a>

<a id="canonical-3200002312310320-2010103010312320-1310112003211310-3213310130330330-3303232022331032-1012101020332013-1111311300013110-3022121322012133"></a>

## max_events property — batch / 130300120132 / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-2303131230201223-3021231221200101-1011333313030310-1312323311331323-1213200313002112-2103321231021313-0302030003033323-1020020312202213): complete subsection reference.

<a id="canonical-3110031303113112-1220203021023112-1322303032012133-2220100302233331-2122203221213033-1132330210231232-1223002221232220-2331101112123130"></a>

<a id="canonical-1011310332220122-3001203211200020-1312120210222023-0310321131301320-3122301220001122-1223301201310210-3020110302333300-3302212131221233"></a>

## timeout_seconds property — batch / 130300120132 / 6

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-3132330132330110-1212133331302311-2011332003332232-1023012032310120-3210201001130002-2113201131021302-2301331000202222-1032011222011113): complete subsection reference.

<a id="canonical-0311233000001000-0012031112230301-3211133310213101-0313311010210231-0002023322102302-3223331100112302-0203030013321312-1231321120231123"></a>

## Next pages — batch / 130300120132 / 7

- [azure_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-001.md#canonical-0303111300332213-1202311101313321-2032103010002210-1001010131030121-3220023313310212-0032232112021333-1232211101312003-3132000132033303)
- [azure_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-2303131230201223-3021231221200101-1011333313030310-1312323311331323-1213200313002112-2103321231021313-0302030003033323-1020020312202213)
- [azure_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-3132330132330110-1212133331302311-2011332003332232-1023012032310120-3210201001130002-2113201131021302-2301331000202222-1032011222011113)
- [azure_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-2121301220330030-1221312033213003-0320321302002303-2133212312111123-3330122101022123-0133210113001032-2201030003210013-3211030023023312)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0303111300332213-1202311101313321-2032103010002210-1001010131030121-3220023313310212-0032232112021333-1232211101312003-3132000132033303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
