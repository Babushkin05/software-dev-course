using System;
namespace sem1
{
	public class Customer
	{
		public string? FIO { get; set; }

		public Car? Car { get; set; }

        public override string ToString()
        {
            if(Car is null)
            {
                return $"{FIO} without car";
            }
            return $"{FIO} with a {Car}";
        }
    }
}

