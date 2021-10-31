create table Orders
(
    OrderUID uuid default uuid_generate(),
    Entry varchar,
    InternalSignature varchar,
    Payment uuid,
    Items uuid,
    Locale varchar,
    CustomerID varchar,
    TrackNumber varchar,
    DeliveryService varchar,
    Shardkey varchar,
    SmID integer
)