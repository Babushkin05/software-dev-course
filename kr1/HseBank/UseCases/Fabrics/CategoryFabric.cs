using HseBank.Domain;

namespace HseBank.UseCases
{
    public static class CategoryFabric
    {
        private static uint maxId = 0;

        public static Category Create(CategoryType categoryType, string name){
            return new Category(maxId++, categoryType, name);
        }
    }
}