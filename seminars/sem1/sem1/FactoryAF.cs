using System;
namespace sem1
{
	public class FactoryAF
	{

		public Stack<Car> Cars { get; private set; }

		public List<Customer> Customers { get; private set; }

        public FactoryAF(List<Customer> customers)
        {
            Customers = customers;

            Cars = new List<Car>();
        }

        public void SaleCar()
        {

        }
    }
}

