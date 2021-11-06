CREATE TABLE public.payment (
                                paymentid uuid NOT NULL,
                                transactioninfo varchar NULL,
                                currency varchar NULL,
                                provider varchar NULL,
                                amount int4 NULL,
                                paymentdt int4 NULL,
                                bank varchar NULL,
                                deliverycost int4 NULL,
                                goodstotal int4 NULL,
                                CONSTRAINT payment_pkey PRIMARY KEY (paymentid)
);

CREATE TABLE public.items (
                              itemid uuid NOT NULL,
                              chrtid int4 NULL,
                              price int4 NULL,
                              rid varchar NULL,
                              "name" varchar NULL,
                              sale int4 NULL,
                              "size" varchar NULL,
                              totalprice int4 NULL,
                              nmid int4 NULL,
                              brand varchar NULL,
                              CONSTRAINT items_pkey PRIMARY KEY (itemid)
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.orders (
                               orderuid uuid default uuid_generate_v4(),
                               entry varchar NULL,
                               internalsignature varchar NULL,
                               payment uuid NOT NULL,
                               items uuid NOT NULL,
                               locale varchar NULL,
                               customerid varchar NULL,
                               tracknumber varchar NULL,
                               deliveryservice varchar NULL,
                               shardkey varchar NULL,
                               smid int4 NULL,
                               CONSTRAINT orders_pkey PRIMARY KEY (orderuid),
                               CONSTRAINT orders_fk1 FOREIGN KEY (items) REFERENCES public.items(itemid),
                               CONSTRAINT orders_fk2 FOREIGN KEY (payment) REFERENCES public.payment(paymentid) ON DELETE CASCADE ON UPDATE CASCADE
);