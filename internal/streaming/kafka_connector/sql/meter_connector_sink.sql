-- Example, currently not in use
CREATE SINK CONNECTOR SINK_METERS_PG WITH (
    'connector.class'                         = 'io.confluent.connect.jdbc.JdbcSinkConnector',
    'connection.url'                          = 'jdbc:postgresql://postgres:5432/postgres',
    'connection.user'                         = 'postgres',
    'connection.password'                     = 'postgres',
    'topics'                                  = 'om_meter_m1,om_meter_m2',
    'key.converter'                           = 'io.confluent.connect.json.JsonSchemaConverter',
    'key.converter.schema.registry.url'       = 'http://schema:8081',
    'value.converter'                         = 'io.confluent.connect.json.JsonSchemaConverter',
    'value.converter.schema.registry.url'     = 'http://schema:8081',
    'auto.create'                             = 'true',
    'auto.evolve'                             = 'true',
    'delete.enabled'                          = 'false',
    'insert.mode'                             = 'upsert',
    'pk.mode'                                 = 'record_key',
    'transforms'                              = 'RenameField,ValueToKey,tsWindowStart,tsWindowEnd',
    'transforms.RenameField.type'             = 'org.apache.kafka.connect.transforms.ReplaceField$Value',
    'transforms.RenameField.renames'          = 'WINDOWSTART_TS:WINDOWSTART,WINDOWEND_TS:WINDOWEND',
    'transforms.ValueToKey.type'              = 'org.apache.kafka.connect.transforms.ValueToKey',
    'transforms.ValueToKey.fields'            = 'WINDOWSTART,WINDOWEND',
    'transforms.tsWindowStart.type'           = 'org.apache.kafka.connect.transforms.TimestampConverter$Key',
    'transforms.tsWindowStart.field'          = 'WINDOWSTART',
    'transforms.tsWindowStart.target.type'    = 'Timestamp',
    'transforms.tsWindowEnd.type'             = 'org.apache.kafka.connect.transforms.TimestampConverter$Key',
    'transforms.tsWindowEnd.field'            = 'WINDOWEND',
    'transforms.tsWindowEnd.target.type'      = 'Timestamp'
);