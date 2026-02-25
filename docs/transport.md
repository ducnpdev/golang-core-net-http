# Test

run: go run server.go
run: go run client.go

- output:
```
ducnp@ip-10-10-181-133 golang-core-net-http % go run cmd/client/main.go 
===== DEFAULT CLIENT =====
error: Get "http://localhost:8080/health": read tcp [::1]:64072->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64075->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64078->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64080->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64079->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64076->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64077->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64081->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64082->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64073->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64083->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64085->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64084->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64086->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64087->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64089->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64088->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64090->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64091->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64092->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64093->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64094->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64095->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64096->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64097->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64099->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64098->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64100->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64101->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64102->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64103->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64104->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64105->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64106->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64107->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64108->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64109->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64110->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64111->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64112->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64113->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64114->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64115->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64116->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64117->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64118->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64119->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64120->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64121->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64122->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64123->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64124->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64125->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64127->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64126->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64129->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64128->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64131->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64130->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64133->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64132->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64137->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64136->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64135->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64134->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64142->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64141->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64138->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64140->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:64139->[::1]:8080: read: connection reset by peer
Total time: 557.81125ms
RPS: 17927.21

===== TUNED CLIENT =====
Total time: 365.806542ms
RPS: 27336.85

===== DEFAULT CLIENT (NO READ BODY) =====
error: Get "http://localhost:8080/health": read tcp [::1]:56759->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56757->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56753->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56781->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56774->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56772->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56780->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56784->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56800->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56798->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56797->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56794->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56795->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:56796->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:58236->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:58235->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:58234->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:58233->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:58237->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59419->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59403->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59402->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59401->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59400->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59399->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59417->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59394->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59363->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59377->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59393->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59375->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59412->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59408->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59405->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59397->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59396->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59357->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59374->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59364->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59382->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59381->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59390->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59370->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59356->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59386->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59413->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59391->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59416->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59387->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59395->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59406->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59383->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59366->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59407->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59389->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59392->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59388->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59368->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59369->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59385->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59367->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59371->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59384->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59376->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59365->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59359->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59379->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59415->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59380->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59378->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59362->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59361->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59360->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59372->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59414->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59410->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59411->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59409->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59358->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59373->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59404->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:59398->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60012->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60011->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60010->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60007->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60008->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60009->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60006->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60005->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60004->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60002->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60000->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60001->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60003->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": write tcp [::1]:61038->[::1]:8080: write: broken pipe
error: Get "http://localhost:8080/health": read tcp [::1]:61039->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61032->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61036->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61026->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61023->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61021->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61018->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61016->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60979->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61007->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60998->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60996->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60995->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60994->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60993->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60992->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61042->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60989->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60991->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60997->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60990->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60988->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60987->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60986->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60985->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60984->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60983->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60982->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60981->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60980->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60978->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:60977->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61041->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62141->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62229->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62243->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62159->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62292->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62291->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62195->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62286->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62189->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62187->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62180->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62181->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62267->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62262->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62255->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62173->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62171->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62148->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62146->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62144->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62254->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62218->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62217->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62213->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62221->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62204->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62202->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62230->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:62232->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:50914->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:50494->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": write tcp [::1]:50951->[::1]:8080: write: broken pipe
error: Get "http://localhost:8080/health": read tcp [::1]:50480->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:50488->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:50477->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:50501->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:50902->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:50502->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:50970->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:50911->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": read tcp [::1]:61202->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61204->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": read tcp [::1]:61076->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": read tcp [::1]:62213->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp 127.0.0.1:51057->127.0.0.1:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp [::1]:61117->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": read tcp [::1]:62286->[::1]:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": dial tcp [::1]:8080: connect: resource temporarily unavailable
error: Get "http://localhost:8080/health": read tcp 127.0.0.1:51055->127.0.0.1:8080: read: connection reset by peer
error: Get "http://localhost:8080/health": read tcp 127.0.0.1:51060->127.0.0.1:8080: read: connection reset by peer
Total time: 1.349238333s
RPS: 7411.59
```