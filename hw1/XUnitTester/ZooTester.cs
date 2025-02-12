namespace XUnitTester;
using Xunit;
using Moq;
using zoo_erp;

public class ZooTests
{
    [Fact]
    /// <summary>
    /// Check getting animal by it id.
    /// </summary>
    public void AddAnimal_Equal_GetItById()
    {
        // Arrang 
        var mockClinic = new Mock<IClinic>();
        mockClinic.Setup(c => c.IsHealthy(It.IsAny<Animal>())).Returns(true);
        Zoo zoo = new Zoo(mockClinic.Object);
        Animal monkey = new Monkey("Lesha", 10, 60, 10);

        // Act 
        int id = zoo.AddAnimal(monkey);

        // Assert 
        Assert.Equal(monkey, zoo.getAnimalById(id));
    }

    [Fact]
    /// <summary>
    /// Check getting thing by id.
    /// </summary>
    public void AddThing_Equal_GetItById()
    {
        // Arrang 
        var mockClinic = new Mock<IClinic>();
        mockClinic.Setup(c => c.IsHealthy(It.IsAny<Animal>())).Returns(true);
        Zoo zoo = new Zoo(mockClinic.Object);
        Thing computer = new Computer("Apple Macbook");

        // Act 
        int id = zoo.AddThing(computer);

        // Assert 
        Assert.Equal(computer, zoo.GetThingById(id));
    }

    [Theory]
    [InlineData(new int[] { 15 })]
    [InlineData(new int[] { 15, 23 })]
    [InlineData(new int[] { 15, 1, 2, 5 })]
    /// <summary>
    /// Check accuracy of counting food consumtion.
    /// </summary>
    /// <param name="foods">list of food consumptions</param>
    public void Check_Copsuntion_Counting(int[] foods)
    {
        // Arrang 
        var mockClinic = new Mock<IClinic>();
        mockClinic.Setup(c => c.IsHealthy(It.IsAny<Animal>())).Returns(true);
        Zoo zoo = new Zoo(mockClinic.Object);

        // Act 
        foreach (int food in foods)
        {
            var mockAnimal = new Mock<Animal>();
            mockAnimal.Setup(t => t.Food).Returns(food);
            zoo.AddAnimal(mockAnimal.Object);
        }

        // Assert 
        Assert.Equal(foods.Sum(), zoo.CountFoodCompsuntion());
    }

    [Fact]
    /// <summary>
    /// Checking returning list of kinds.
    /// </summary>
    public void Add2KindAnimals_Return2KindAnimals()
    {
        // Arrange 
        var mockClinic = new Mock<IClinic>();
        mockClinic.Setup(c => c.IsHealthy(It.IsAny<Animal>())).Returns(true);
        Zoo zoo = new Zoo(mockClinic.Object);
        Animal rabbit1 = new Rabbit("Name", 10, 100, 10);
        Animal rabbit2 = new Rabbit("Name", 10, 100, 8);
        Animal rabbit3 = new Rabbit("Name", 10, 100, 5);
        Animal wolf = new Wolf("Name", 10, 100);

        // Act 
        zoo.AddAnimal(rabbit1);
        zoo.AddAnimal(rabbit2);
        zoo.AddAnimal(rabbit3);
        zoo.AddAnimal(wolf);

        // Assert 
        Assert.Equal(2, zoo.getKindAnimals().Count());
    }

    [Fact]
    /// <summary>
    /// Check not adding ill animals.
    /// </summary>
    public void CheckUnability_ToAddIllAnimals()
    {
        // Arrange 
        var mockClinic = new Mock<IClinic>();
        mockClinic.Setup(c => c.IsHealthy(It.IsAny<Animal>())).Returns(false);
        Zoo zoo = new Zoo(mockClinic.Object);

        // Act 
        int id = zoo.AddAnimal(new Wolf("Vova", 10, 100));

        // Assert 
        Assert.Equal(-1, id);
    }
}