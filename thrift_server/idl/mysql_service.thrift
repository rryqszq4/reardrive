namespace go mysql_service

struct MysqlStruct {
    1: string   host
    2: i32      port
    3: string   user
    4: string   password
    5: string   database
    6: string   charset
}

struct QueryStruct {
    1: string   sql
}

struct ItemStruct {
    1: map<string,string> data
}

struct QueryResultStruct {
    1: list<ItemStruct>     rows
    2: string               error
    3: QueryStruct          query
    4: map<string,string>   columns
}

struct ExecuteResultStruct {
    1: string       result
    2: string       error
    3: i64          lastInsertId
    4: i64          rowsAffected
    5: QueryStruct  query
}

service MysqlService {
    ExecuteResultStruct     execute(1:MysqlStruct mysql, 2:QueryStruct query)

    QueryResultStruct       query(1:MysqlStruct mysql, 2: QueryStruct query)
}