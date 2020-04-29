export default interface Event { 
	ID: String;
	name?: String;
	start_date: Date;
	end_date: Date;
	location: {
		ID?: String;
		name?: String;
		country?: String;
	};
}