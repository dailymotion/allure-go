import io.qameta.allure.Description;
import org.testng.annotations.Parameters;
import org.testng.annotations.Test;

public class AllureBasicTest {
    @Test
    public void test() {
        StepExamples.doStuff();
    }

    @Test
    @Description("test description")
    @Parameters({ "xmlPath" })
    public void parameterTest(String parameter) {
        System.out.println(parameter);
        StepExamples.doStuff();
    }
}