import io.qameta.allure.Step;

import static io.qameta.allure.Allure.step;

public class StepExamples {

    @Step("Other stuff")
    public static void doOtherStuff() {
    }

    @Step("Do stuff")
    public static void doStuff() {
        doOtherStuff();
        System.out.println("");
//        step("step 1");
//        step("step 2");
    }
}
