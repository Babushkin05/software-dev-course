using System;
namespace sem1
{
	public class FactoryAF
	{

		public Stack<Car> Cars { get; private set; }

		public List<Customer> Customers { get; private set; }

        public FactoryAF(params Customer[] customers)
        {
            Customers = new List<Customer>(customers);

            Cars = new Stack<Car>();
        }


        public void SaleCar()
        {

            for(int i = 0; i<Customers.Count(); ++i)
            {
                if(Cars.Count == 0)
                {
                    return;
                }
                Customers[i].Car = Cars.Pop();
                if(i == Customers.Count() - 1)
                {
                    Cars.Clear();
                }
            }
        }

    }
}

