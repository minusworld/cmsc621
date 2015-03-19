package dict3

func setBooleanReply(reply *bool, err error) {
	if err != nil {
		*reply = false
	} else {
		*reply = true
	}
}

func (d *DICT3) RLookup(args KeyRelationship, reply *interface{}) error {
	value, err := d.Lookup(args.Key, args.Relationship)
	*reply = value
	return err
}

func (d *DICT3) RInsert(args KeyRelationshipValue, reply *bool) error {
	err := d.Insert(args.Key, args.Relationship, args.Value)
	setBooleanReply(reply, err)
	return err
}

func (d *DICT3) RInsertOrUpdate(args KeyRelationshipValue, reply *bool) error {
	err := d.InsertOrUpdate(args.Key, args.Relationship, args.Value)
	setBooleanReply(reply, err)
	return err
}

func (d *DICT3) RDelete(args KeyRelationship, reply *bool) error {
	err := d.Delete(args.Key, args.Relationship)
	setBooleanReply(reply, err)
	return err
}

func (d *DICT3) RListKeys(args struct{}, reply *[]string) error {
	keys, err := d.ListKeys()
	*reply = keys
	return err
}

func (d *DICT3) RListIDs(args struct{}, reply *[]string) error {
	keys, err := d.ListIDs()
	*reply = keys
	return err
}

func (d *DICT3) RShutdown(args struct{}, reply *bool) error {
	defer Shutdown()
	*reply = true
	return nil
}